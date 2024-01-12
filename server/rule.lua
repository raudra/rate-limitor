--ngx.say('Hello,world!')
local json = require 'json'
local Rule={window=0, count=0, name=""}
local RateLimitor={rules={}, ipAddress="", path="", cache=nil, rules = {}}

function Rule:new(o, window, count, name)
	o = o or {}
	setmetatable(o, self)
	self.__index = self
	o.window = window
	o.count = count
	o.name = name
	return o
end

function Rule:buildKey(ipAddress, path)
	if self.name == 'IpRule' then
		return "IP::" .. ipAddress

	elseif self.name == 'IPPathRule' then
		return "Path::" .. path .. "::IP::" ..ipAddress 
	else 
		print('Missing rule defination -- ', self)
	end
end


function RateLimitor:new(o, rules, ipAddress, path)
	local redis = require 'resty.redis'
	o = o or {}
	setmetatable(o, self)
	self.__index = self
	o.rules = rules
	o.ipAddress = ipAddress
	o.path = path
	o.cache = RedisClient(redis)
	return o
end

function Rule:__tostring()
	print("window: ", self.window, " count: ", self.count, " name: ", self.name)
end

function RedisClient(redis)
	local red = redis:new()
	local options_table = {}
	options_table["pool"] = "docker_server"
	local ok, err = red:connect('192.168.0.103', 6300, options_table) 
	
	if not ok then
                    ngx.say("failed to connect: ", err)
                    return
        end

	return red
end

function RateLimitor:getRules()
	local rules ={}

	if self.cache == nil then
		ngx.say("Cache should be initialize")
		return rules
	else
		local array, ok = self.cache:smembers("Rules")
		
		for k, v in pairs(array) do
			local data = v
			if data ~= nil then
				local tbl = json.decode(data)
				local r = Rule:new(nil, tbl["window"], tbl["count"], tbl["name"])	
				table.insert(rules,r)
			end

		end
		return rules
	end	
end


function RateLimitor:validate()
	local flag=true
	self.rules = self:getRules()
	for k, v in pairs(self.rules) do
		if v ~= nil then
			flag = self:triggerAlgo(v)
		end
		if not flag then
			ngx.exit(422)
			break
		end
	end
end

function RateLimitor:triggerAlgo(rule)
	local key = rule:buildKey(self.ipAddress, self.path )
	local window = rule.window
	local count  = rule.count
	
	if key == nil then 
		return false
	end

	local currentTime = self.cache:time()
	local trimTime = tonumber(currentTime[1]) -  window
	self.cache:zremrangebyscore(key, 0, trimTime)

	local reqCount = table.getn(self.cache:zrange(key, 0, -1))
	
	print("Request Count " .. reqCount .. "Limit" ..count.. " Key " ..key)
	if reqCount < count then
		self.cache:zadd(key, currentTime[1], currentTime[1] .. currentTime[2])
		self.cache:expire(key, window)
		return true
	end
	return false
end

ngx.req.read_body()

local rl = RateLimitor:new(nil, rules, ngx.var.remote_addr, ngx.var.request_uri)
rl:validate()


