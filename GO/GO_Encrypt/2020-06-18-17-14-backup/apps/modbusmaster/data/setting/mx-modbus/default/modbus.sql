PRAGMA user_version=1;

-- TABLE
CREATE TABLE [config](
	[configId] INTEGER PRIMARY KEY
);
CREATE TABLE [master-params](
	[masterParamId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL DEFAULT 1,
	[enableTcpMasters] INT NOT NULL DEFAULT 1,
	[enableSerMasters] INT NOT NULL DEFAULT 1,
	[enableDevFailEvent] INT NOT NULL DEFAULT 1,
	[enableCmdFailEvent] INT NOT NULL DEFAULT 1,
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [enableTcpMasters (0,1)] CHECK ([enableTcpMasters] IN (0,1)),
	CONSTRAINT [enableSerMasters (0,1)] CHECK ([enableSerMasters] IN (0,1)),
	CONSTRAINT [enableDevFailEvent (0,1)] CHECK ([enableDevFailEvent] IN (0,1)),
	CONSTRAINT [enableCmdFailEvent (0,1)] CHECK ([enableCmdFailEvent] IN (0,1))
);
CREATE TABLE [master-ser-ifaces](
	[masterSerIfaceId] INTEGER PRIMARY KEY,
	[serMasterId] INTEGER NOT NULL,
	[portValue] INT NOT NULL,
	[format] INT NOT NULL DEFAULT 0,
	[initialDelay] INT NOT NULL DEFAULT 0,
	[retryCount] INT NOT NULL DEFAULT 3,
	[responseTout] INT NOT NULL DEFAULT 1000,
	[frameInterval] INT NOT NULL DEFAULT 0,
	[charInterval] INT NOT NULL DEFAULT 0,
	CONSTRAINT [fkSerMaster] FOREIGN KEY ([serMasterId]) REFERENCES [ser-masters]([serMasterId]) ON DELETE CASCADE ON UPDATE CASCADE,
	UNIQUE (portValue),
	CONSTRAINT [format (0,1)] CHECK ([format] IN (0,1)),
	CONSTRAINT [initialDelay (0-30000)] CHECK ([initialDelay] BETWEEN 0 AND 30000),
	CONSTRAINT [retryCount (0-5)] CHECK ([retryCount] BETWEEN 0 AND 5),
	CONSTRAINT [responseTout (10-120000)] CHECK ([responseTout] BETWEEN 10 AND 120000),
	CONSTRAINT [frameInterval (0,10-500)] CHECK ([frameInterval] IN (0) OR ([frameInterval] BETWEEN 10 AND 500)),
	CONSTRAINT [charInterval (0, 10-500)] CHECK ([charInterval] IN (0) OR ([charInterval] BETWEEN 10 AND 500))
);
CREATE TABLE [master-tcp-ifaces](
	[masterTcpIfaceId] INTEGER PRIMARY KEY,
	[tcpMasterId] INTEGER NOT NULL,
	[initialDelay] INT NOT NULL DEFAULT 0,
	[retryCount] INT NOT NULL DEFAULT 3,
	[responseTout] INT NOT NULL DEFAULT 1000,
	CONSTRAINT [fkTcpMaster] FOREIGN KEY ([tcpMasterId]) REFERENCES [tcp-masters]([tcpMasterId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [initialDelay (0-30000)] CHECK ([initialDelay] BETWEEN 0 AND 30000),
	CONSTRAINT [retryCount (0-5)] CHECK ([retryCount] BETWEEN 0 AND 5),
	CONSTRAINT [responseTout (10-120000)] CHECK ([responseTout] BETWEEN 10 AND 120000)
);
CREATE TABLE [mcmds](
	[mcmdId] INTEGER PRIMARY KEY, 
	[remoteDevId] INTEGER NOT NULL,
	[name] TEXT NOT NULL DEFAULT "command", 
	[enable] INT NOT NULL DEFAULT 1,
	[mode] INT NOT NULL DEFAULT 0,  
	[func] INT NOT NULL DEFAULT 3,  
	[readAddress] INT NOT NULL DEFAULT 0,
	[readQuantity] INT NOT NULL DEFAULT 10,
	[writeAddress] INT NOT NULL DEFAULT 0,  
	[writeQuantity] INT NOT NULL DEFAULT 1,  
	[pollInterval] INT NOT NULL DEFAULT 1000,
	[swap] INT NOT NULL DEFAULT 0,
	[fpFunc] INT NOT NULL DEFAULT 0,
	[fpTout] INT NOT NULL DEFAULT 3600,
	[fpData] TEXT NOT NULL DEFAULT "00 00",
	[scalingFunc] INT NOT NULL DEFAULT 0,
	[interceptSlope] REAL NOT NULL DEFAULT 1,
	[interceptOffset] REAL NOT NULL DEFAULT 0,
	[pointSourceMin] REAL NOT NULL DEFAULT 0,
	[pointSourceMax] REAL NOT NULL DEFAULT 1,
	[pointTargetMin] REAL NOT NULL DEFAULT 0,
	[pointTargetMax] REAL NOT NULL DEFAULT 1,
	CONSTRAINT [fkRemoteDev] FOREIGN KEY ([remoteDevId]) REFERENCES [remote-devs]([remoteDevId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [name length (1-120)] CHECK (length(name) BETWEEN 1 AND 120),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	UNIQUE (remoteDevId, name),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1)),
	CONSTRAINT [mode (0,1)] CHECK ([mode] IN (0,1)),
	CONSTRAINT [func (1, 2, 3, 4, 5, 6, 15, 16, 23)] CHECK ([func] IN (1, 2, 3, 4, 5, 6, 15, 16, 23)),
	CONSTRAINT [readAddress (0-65535) & (readAddress + readQuantity <= 65536)] CHECK (([readAddress] BETWEEN 0 AND 65535) AND ([readAddress] + [readQuantity] <= 65536)),
	CONSTRAINT [readQuantity >= 1] CHECK ([readQuantity] >= 1),
	CONSTRAINT [(func (1,2) & readQuantity <= 2000) or (func (3,4,23) & readQuantity <= 125)] CHECK (([func] IN (1,2) AND [readQuantity] <= 2000) OR ([func] IN (3,4,23) AND [readQuantity] <= 125) OR ([func] IN (5,6,15,16))),	
	CONSTRAINT [writeAddress (0-65535) & (writeAddress + writeQuantity <= 65536)] CHECK (([writeAddress] BETWEEN 0 AND 65535) AND ([writeAddress] + [writeQuantity] <= 65536)),
	CONSTRAINT [writeQuantity >= 1] CHECK ([writeQuantity] >= 1),
	CONSTRAINT [(func (1,2,3,4)) or (func (5,6) & writeQuantity <= 1) or (func (15) & writeQuantity <= 1968) or (func (16) & writeQuantity <= 123) or (func (23) & writeQuantity <= 121)] CHECK (([func] IN (1,2,3,4)) OR ([func] IN (5,6) AND [writeQuantity] <= 1) OR ([func] IN (15) AND [writeQuantity] <= 1968) OR ([func] IN (16) AND [writeQuantity] <= 123) OR ([func] IN (23) AND [writeQuantity] <= 121)),
	CONSTRAINT [pollInterval (100-86400000)] CHECK ([pollInterval] BETWEEN 100 AND 86400000),
	CONSTRAINT [swap (0,1,2,3)] CHECK ([swap] IN (0,1,2,3)),
	CONSTRAINT [fpFunc (0,1,3)] CHECK ([fpFunc] IN (0,1,3)),
	CONSTRAINT [fpTout (1-86400)] CHECK ([fpTout] BETWEEN 1 AND 86400),
	CONSTRAINT [scalingFunc (0,1,2)] CHECK ([scalingFunc] IN (0,1,2))
);
CREATE TABLE [ref-tags](
	[refTagId] INTEGER PRIMARY KEY,
	[szoneId] INTEGER NOT NULL,
	[prvdName] TEXT NOT NULL,
	[srcName] TEXT NOT NULL,
	[tagName] TEXT NOT NULL,
	[access] TEXT NOT NULL,	
	[dataType] TEXT NOT NULL,
	[dataSize] INT NULL,
	CONSTRAINT [fkSzone] FOREIGN KEY ([szoneId]) REFERENCES [szones]([szoneId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [prvdName length (1-128)] CHECK (length(prvdName) BETWEEN 1 AND 128),
	CONSTRAINT [prvdName only accept - . _ ~ 0-9 a-z A-Z] CHECK ([prvdName] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [srcName length (1-128)] CHECK (length(srcName) BETWEEN 1 AND 128),
	CONSTRAINT [srcName only accept - . _ ~ 0-9 a-z A-Z] CHECK ([srcName] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [tagName length (1-128)] CHECK (length(tagName) BETWEEN 1 AND 128),
	CONSTRAINT [tagName only accept - . _ ~ 0-9 a-z A-Z] CHECK ([tagName] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [access (r,w,rw)] CHECK ([access] IN ("r","w","rw")),
	CONSTRAINT [dataType (boolean, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float, double, string, bytearray, raw)] CHECK ([dataType] IN ("boolean", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float", "double", "string", "bytearray", "raw")),
	CONSTRAINT [dataSize should exist when dataType is (string, bytearray, raw), otherwise it should not exist] CHECK (([dataType] NOT IN ("string", "bytearray", "raw")) OR ([dataType] IN ("string", "bytearray", "raw") AND ([dataSize] IS NOT NULL) AND ([dataSize] >= 1)))
);
CREATE TABLE [remote-devs](
	[remoteDevId] INTEGER PRIMARY KEY,
	[masterSerIfaceId] INTEGER NULL,  
	[masterTcpIfaceId] INTEGER NULL,  
	[name] TEXT UNIQUE NOT NULL DEFAULT "device",  
	[enable] INT NOT NULL DEFAULT 1,
	[slaveId] INT NOT NULL DEFAULT 1,
	[slaveIpaddr] TEXT NOT NULL DEFAULT "0.0.0.0",
	[slaveTcpPort] INT NOT NULL DEFAULT 502,
	CONSTRAINT [fkmasterSerIfacesDev] FOREIGN KEY ([masterSerIfaceId]) REFERENCES [master-ser-ifaces]([masterSerIfaceId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [fkmasterTcpIfacesDev] FOREIGN KEY ([masterTcpIfaceId]) REFERENCES [master-tcp-ifaces]([masterTcpIfaceId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [exactly one of masterSerIfaceId or masterTcpIfaceId] CHECK ((([masterSerIfaceId] IS NULL) OR ([masterTcpIfaceId] IS NULL)) AND (([masterSerIfaceId] IS NOT NULL) OR ([masterTcpIfaceId] IS NOT NULL))),
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1)),
	CONSTRAINT [slaveId (0-255)] CHECK ([slaveId] BETWEEN 0 AND 255),
	CONSTRAINT [slaveTcpPort (1-65535)] CHECK ([slaveTcpPort] BETWEEN 1 AND 65535),
	UNIQUE ([masterSerIfaceId], [slaveId]),
	UNIQUE ([masterTcpIfaceId], [slaveId], [slaveIpaddr], [slaveTcpPort])
);
CREATE TABLE [ser-masters](
	[serMasterId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL,
	[name] TEXT UNIQUE NOT NULL DEFAULT "modbus_serial_master",
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*')
);
CREATE TABLE [ser-slaves](
	[serSlaveId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL,
	[slaveId] INT UNIQUE NOT NULL DEFAULT 1,
	[format] INT NOT NULL DEFAULT 0,
	[name] TEXT UNIQUE NOT NULL DEFAULT "modbus_serial_slave",
	[enable] INT NOT NULL DEFAULT 0,
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [slaveId (1-255)] CHECK ([slaveId] BETWEEN 1 AND 255),
	CONSTRAINT [format (0,1)] CHECK ([format] IN (0,1)),
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1))
);
CREATE TABLE [slave-params](
	[slaveParamId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL DEFAULT 1,
	[enableTcpSlaves] INT NOT NULL DEFAULT 1,
	[enableSerSlaves] INT NOT NULL DEFAULT 1,
	[enableFailEvent] INT NOT NULL DEFAULT 1,
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [enableTcpSlaves (0,1)] CHECK ([enableTcpSlaves] IN (0,1)),
	CONSTRAINT [enableSerSlaves (0,1)] CHECK ([enableSerSlaves] IN (0,1)),
	CONSTRAINT [enableFailEvent (0,1)] CHECK ([enableFailEvent] IN (0,1))
);
CREATE TABLE [slave-ser-ifaces](
	[slaveSerIfaceId] INTEGER PRIMARY KEY,
	[serSlaveId] INTEGER NOT NULL,
	[portValue] INT UNIQUE NOT NULL,
	CONSTRAINT [fkSerSlave] FOREIGN KEY ([serSlaveId]) REFERENCES [ser-slaves]([serSlaveId]) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE [szones](
	[szoneId] INTEGER PRIMARY KEY, 
	[serSlaveId] INTEGER NULL,  
	[tcpSlaveId] INTEGER NULL,
	[dataFormat] INT NOT NULL DEFAULT 2,  
	[address] INT NOT NULL DEFAULT 0,
	[quan] INT NOT NULL DEFAULT 1,
	[fromRefTaglist] INT NOT NULL DEFAULT 0,
	CONSTRAINT [fkSerSlave] FOREIGN KEY ([serSlaveId]) REFERENCES [ser-slaves]([serSlaveId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [fkTcpSlave] FOREIGN KEY ([tcpSlaveId]) REFERENCES [tcp-slaves]([tcpSlaveId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [exactly one of serSlaveId or tcpSlaveId] CHECK ((([serSlaveId] IS NULL) OR ([tcpSlaveId] IS NULL)) AND (([serSlaveId] IS NOT NULL) OR ([tcpSlaveId] IS NOT NULL))),
	CONSTRAINT [dataFormat (0,1,2,3)] CHECK ([dataFormat] IN (0,1,2,3)),
	CONSTRAINT [address (0-65535)] CHECK ([address] BETWEEN 0 AND 65535),
	CONSTRAINT [quan >= 1 & (quan + address <= 65536)] CHECK (([quan] >= 1) AND ([quan] + [address] <= 65536)),
	CONSTRAINT [fromRefTaglist (0,1)] CHECK ([fromRefTaglist] IN (0,1))
	
);
CREATE TABLE [tags](
	[tagId] INTEGER PRIMARY KEY,
	[mcmdId] INTEGER NULL,
	[szoneId] INTEGER NULL,
	[name] TEXT NOT NULL DEFAULT 'tag',
	[dataType] TEXT NOT NULL DEFAULT 'raw',
	[dataUnit] TEXT NULL,
	[access] TEXT NOT NULL DEFAULT 'rw',
	[dataSize] INT NULL,
	[offset] INT NOT NULL DEFAULT 0,
	CONSTRAINT [fkMcmd] FOREIGN KEY ([mcmdId]) REFERENCES [mcmds]([mcmdId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [fkSzone] FOREIGN KEY ([szoneId]) REFERENCES [szones]([szoneId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [exactly one of mcmdId or szoneId] CHECK ((([mcmdId] IS NULL) OR ([szoneId] IS NULL)) AND (([mcmdId] IS NOT NULL) OR ([szoneId] IS NOT NULL))),
	CONSTRAINT [name length (1,128)] CHECK (length(name) BETWEEN 1 AND 128),	
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [name should not be (status)] CHECK ([name] != 'status'),
	CONSTRAINT [dataType (boolean, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float, double, string, bytearray, raw)] CHECK ([dataType] IN ("boolean", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float", "double", "string", "bytearray", "raw")),
	CONSTRAINT [access (r,w,rw)] CHECK ([access] IN ("r","w","rw")),
	CONSTRAINT [dataSize should exist when dataType is (string, bytearray, raw), otherwise it should not exist] CHECK (([dataType] NOT IN ("string", "bytearray", "raw")) OR ([dataType] IN ("string", "bytearray", "raw")  AND ([dataSize] IS NOT NULL) AND ([dataSize] >= 1)))
);
CREATE TABLE [tcp-masters](
	[tcpMasterId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL,
	[name] TEXT UNIQUE NOT NULL DEFAULT "modbus_tcp_master",
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*')
);
CREATE TABLE [tcp-slaves](
	[tcpSlaveId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL,
	[slaveId] INT NOT NULL DEFAULT 1,
	[listenPort] INT NOT NULL DEFAULT 502,
	[name] TEXT UNIQUE NOT NULL DEFAULT "modbus_tcp_slave",
	[enable] INT NOT NULL DEFAULT 0,
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	UNIQUE ([slaveId], [listenPort]),
	CONSTRAINT [slaveId (1-255)] CHECK ([slaveId] BETWEEN 1 AND 255),
	CONSTRAINT [listenPort (1-65535)] CHECK ([listenPort] BETWEEN 1 AND 65535),
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1))
);
 
-- INDEX
 
CREATE INDEX [idxTags] ON [tags] (mcmdId);

-- TRIGGER
CREATE TRIGGER chkTopicInsert
BEFORE INSERT ON [tags]
WHEN 
	(
		(
			(NEW.[mcmdId] IS NOT NULL) 
			AND
			(			
				SELECT COUNT(*)
				FROM (
					SELECT [mcmds].remoteDevId AS refDevId FROM [tags]
					INNER JOIN [mcmds] ON [tags].name == NEW.[name] AND [mcmds].mcmdId == [tags].mcmdId 
					) AS filterTable
				WHERE (SELECT [mcmds].remoteDevId FROM [mcmds] WHERE [mcmds].mcmdId == NEW.mcmdId AND [mcmds].remoteDevId == filterTable.refDevId)
			)
		)
		OR
		(
			(NEW.[szoneId] IS NOT NULL)
			AND
			(
				SELECT COUNT(*) FROM [tags]
				INNER JOIN [szones] ON [szones].szoneId == [tags].szoneId AND [tags].name == NEW.[name]				
			)
		)
	) > 0
BEGIN
    SELECT RAISE(ABORT, 'The set of following column(s) should be unique: remoteDevId, tags.name');
END;
CREATE TRIGGER chkTopicUpdate
BEFORE UPDATE ON [tags]
WHEN 
	(
		(
			(NEW.[mcmdId] IS NOT NULL) 
			AND
			(
				SELECT COUNT(*)
				FROM (
					SELECT [mcmds].remoteDevId AS refDevId FROM [tags]
					INNER JOIN [mcmds] ON [tags].name == NEW.[name] AND OLD.[name] != NEW.[name] AND [mcmds].mcmdId == [tags].mcmdId 
					) AS filterTable
				WHERE (SELECT [mcmds].remoteDevId FROM [mcmds] WHERE [mcmds].mcmdId == NEW.mcmdId AND [mcmds].remoteDevId == filterTable.refDevId)
			)
		)
		OR
		(
			(NEW.[szoneId] IS NOT NULL)
			AND
			(
				SELECT COUNT(*) FROM [tags]
				INNER JOIN [szones] ON [szones].szoneId == [tags].szoneId AND [tags].name == NEW.[name]				
			)
		)
	) > 0
BEGIN
    SELECT RAISE(ABORT, 'remoteDevId, tags.name exists already');
END;
CREATE TRIGGER [chkMaxConfigCount] 
BEFORE INSERT ON [config]
WHEN (
    SELECT COUNT(*) FROM [config] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "config are full(1)");
END;

CREATE TRIGGER [chkMaxMasterParamsCount] 
BEFORE INSERT ON [master-params]
WHEN (
    SELECT COUNT(*) FROM [master-params] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "master-params are full(1)");
END;

CREATE TRIGGER [chkMaxMasterSerIfaceCount] 
BEFORE INSERT ON [master-ser-ifaces]
WHEN (
    SELECT COUNT(*) FROM [master-ser-ifaces] 
    ) >= @PORT_NUM@
BEGIN
    SELECT RAISE(ABORT, "master-ser-ifaces are full(@PORT_NUM@)");
END;

CREATE TRIGGER [chkMaxMasterTcpIfacesCount] 
BEFORE INSERT ON [master-tcp-ifaces]
WHEN (
    SELECT COUNT(*) FROM [master-tcp-ifaces] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "master-tcp-ifaces are full(1)");
END;
CREATE TRIGGER [chkMaxMcmdCount]
BEFORE INSERT ON [mcmds]
WHEN (
    SELECT COUNT (*) FROM [mcmds]
)	>= 2048
BEGIN
    SELECT RAISE(ABORT, 'mcmds are full(2048)');
END;
CREATE TRIGGER [chkMaxMcmdCountInOneSerPort] 
BEFORE INSERT ON [mcmds]
WHEN (
    SELECT COUNT(*) FROM [mcmds] 
    	INNER JOIN [remote-devs] ON 
        	[remote-devs].remoteDevId = [mcmds].remoteDevId AND
            [mcmds].remoteDevId == NEW.[remoteDevId]
        INNER JOIN [master-ser-ifaces] ON
        	[master-ser-ifaces].masterSerIfaceId == [remote-devs].masterSerIfaceId
    ) >= 128
BEGIN
    SELECT RAISE(ABORT, "mcmds in this serial port are full(128)");
END;
CREATE TRIGGER [chkMaxMcmdCountInOneTcpPort] BEFORE INSERT ON [mcmds]
WHEN (
	SELECT COUNT(*) FROM [mcmds] 
    	INNER JOIN [remote-devs] ON 
        	[remote-devs].remoteDevId = [mcmds].remoteDevId AND
            [mcmds].remoteDevId == NEW.[remoteDevId]
        INNER JOIN [master-tcp-ifaces] ON
        	[master-tcp-ifaces].masterTcpIfaceId == [remote-devs].masterTcpIfaceId
    ) >= 2048
BEGIN
    SELECT RAISE(ABORT, "mcmds in this TCP port are full(2048)");
END;
CREATE TRIGGER [chkMaxRefTagsCount]
BEFORE INSERT ON [ref-tags]
WHEN (
    SELECT COUNT (*) FROM [ref-tags]
)	>= 2048
BEGIN
    SELECT RAISE(ABORT, 'ref-tags are full(2048)');
END;
CREATE TRIGGER [chkMaxSerMastersCount] 
BEFORE INSERT ON [ser-masters]
WHEN (
    SELECT COUNT(*) FROM [ser-masters] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "ser-masters are full(1)");
END;

CREATE TRIGGER [chkMaxSerSlavesCount] 
BEFORE INSERT ON [ser-slaves]
WHEN (
    SELECT COUNT(*) FROM [ser-slaves] 
    ) >= @PORT_NUM@
BEGIN
    SELECT RAISE(ABORT, "ser-slaves are full(@PORT_NUM@)");
END;
CREATE TRIGGER [chkMaxSlaveSerIfacesCount] 
BEFORE INSERT ON [slave-ser-ifaces]
WHEN (
    SELECT COUNT(*) FROM [slave-ser-ifaces] 
    ) >= @PORT_NUM@
BEGIN
    SELECT RAISE(ABORT, "slave-ser-ifaces are full(@PORT_NUM@)");
END;
CREATE TRIGGER [chkMaxSerRemoteDevsCount]
BEFORE INSERT ON [remote-devs]
WHEN(
    (NEW.[masterSerIfaceId] NOT NULL) AND ((SELECT COUNT (*) FROM [remote-devs] WHERE [masterSerIfaceId] NOT NULL) >= 31)
	)
BEGIN
    SELECT RAISE(ABORT, 'Serial remote-devs are full(31)');
END;
CREATE TRIGGER [chkMaxSlaveParamsCount] 
BEFORE INSERT ON [slave-params]
WHEN (
    SELECT COUNT(*) FROM [slave-params] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "slave-params are full(1)");
END;
CREATE TRIGGER [chkMaxSzonesCount] 
BEFORE INSERT ON [szones]
WHEN (
    SELECT COUNT(*) FROM [szones] 
    ) >= 2048
BEGIN
    SELECT RAISE(ABORT, "szones are full(2048)");
END;
CREATE TRIGGER [chkMaxTagInAllMasterCount]
BEFORE INSERT ON [tags]
WHEN 
	(
			(NEW.[mcmdId] IS NOT NULL) 
			AND
			(
				(SELECT COUNT(*) FROM [tags] WHERE mcmdId IS NOT NULL) >= 2048
			)
	) 
BEGIN
    SELECT RAISE(ABORT, 'Master tags are full(2048)');
END;
CREATE TRIGGER [chkMaxTagInAllSlaveCount]
BEFORE INSERT ON [tags]
WHEN 
	(
			(NEW.[szoneId] IS NOT NULL) 
			AND
			(
				(SELECT COUNT(*) FROM [tags] WHERE szoneId IS NOT NULL) >= 2048
			)
	) 
BEGIN
    SELECT RAISE(ABORT, 'Slave tags are full(2048)');
END;
CREATE TRIGGER [chkMaxTcpMastersCount] 
BEFORE INSERT ON [tcp-masters]
WHEN (
    SELECT COUNT(*) FROM [tcp-masters] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "tcp-masters are full(1)");
END;
CREATE TRIGGER [chkMaxTcpRemoteDevsCount]
BEFORE INSERT ON [remote-devs]
WHEN(
	 (NEW.[masterTcpIfaceId] NOT NULL) AND ((SELECT COUNT (*) FROM [remote-devs] WHERE [masterTcpIfaceId] NOT NULL) >= 32)
	 )
BEGIN
    SELECT RAISE(ABORT, 'TCP remote-devs are full(32)');
END;
CREATE TRIGGER [chkMaxTcpSlavesCount] 
BEFORE INSERT ON [tcp-slaves]
WHEN (
    SELECT COUNT(*) FROM [tcp-slaves] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "tcp-slaves are full(1)");
END;
CREATE TRIGGER chkRefTagsUniqueNameInsert
BEFORE INSERT ON [ref-tags]
FOR EACH ROW
WHEN ((SELECT refTagId FROM [ref-tags] WHERE [ref-tags].prvdName = NEW.prvdName AND [ref-tags].srcName = NEW.srcName AND [ref-tags].tagName = NEW.tagName) IS NOT NULL)
BEGIN
    SELECT RAISE(ABORT, 'ref-tags.prvdName, ref-tags.srcName, ref-tags.tagName exists already');
END;
CREATE TRIGGER [chkRefTagsUniqueNameUpdate]
BEFORE UPDATE ON [ref-tags]
FOR EACH ROW
WHEN ((SELECT refTagId FROM [ref-tags] WHERE [ref-tags].prvdName = NEW.prvdName AND [ref-tags].srcName = NEW.srcName AND [ref-tags].tagName = NEW.tagName AND OLD.tagName != NEW.tagName) IS NOT NULL)
BEGIN
    SELECT RAISE(ABORT, 'ref-tags.prvdName, ref-tags.srcName, ref-tags.tagName exists already');
END;
 
CREATE TRIGGER setTagsDataSizeNullInsert
AFTER INSERT ON [tags]
WHEN (NEW.dataType NOT IN ('raw','bytearray','string'))
BEGIN
   UPDATE tags SET dataSize = null WHERE ROWID = NEW.ROWID;
END;

CREATE TRIGGER setTagsDataSizeNullUpdate
AFTER UPDATE ON [tags]
WHEN (NEW.dataType NOT IN ('raw','bytearray','string'))
BEGIN
   UPDATE tags SET dataSize = null WHERE ROWID = NEW.ROWID;
END;

CREATE TRIGGER setReftagsDataSizeNullInsert
AFTER INSERT ON [ref-tags]
WHEN (NEW.dataType NOT IN ('raw','bytearray','string'))
BEGIN
   UPDATE [ref-tags] SET dataSize = null WHERE ROWID = NEW.ROWID;
END;

CREATE TRIGGER setReftagsDataSizeNullUpdate
AFTER UPDATE ON [ref-tags]
WHEN (NEW.dataType NOT IN ('raw','bytearray','string'))
BEGIN
   UPDATE [ref-tags] SET dataSize = null WHERE ROWID = NEW.ROWID;
END;

-- VIEW
 
