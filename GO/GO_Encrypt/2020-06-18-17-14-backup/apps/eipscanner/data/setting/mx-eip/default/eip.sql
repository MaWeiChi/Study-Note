PRAGMA user_version=1;

-- TABLE
CREATE TABLE [adapters](
	[adapterId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL,
	[outputSize] INT NOT NULL DEFAULT 0,
	[inputSize] INT NOT NULL DEFAULT 0,
	[eipEncapInactivityTimeoutValue] INT NOT NULL DEFAULT 120,
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [outputSize (0-496)] CHECK ([outputSize] BETWEEN 0 AND 496),
	CONSTRAINT [inputSize (0-496)] CHECK ([inputSize] BETWEEN 0 AND 496),	
	CONSTRAINT [eipEncapInactivityTimeoutValue (0-3600)] CHECK ([eipEncapInactivityTimeoutValue] BETWEEN 0 AND 3600)
);
CREATE TABLE [config](
	[configId] INTEGER PRIMARY KEY,
	[eipAssemblyMode] BOOLEAN NOT NULL DEFAULT 0, 
	CONSTRAINT [eipAssemblyMode (0,1)] CHECK ([eipAssemblyMode] IN (0,1)) /* 0 = adapter, 1 = scanner */	
);
CREATE TABLE [connections](
	[connectionId] INTEGER PRIMARY KEY, 
	[deviceId] INTEGER NOT NULL,
	[name] TEXT NOT NULL DEFAULT "connection", 
	[enable] BOOLEAN NOT NULL DEFAULT 1,
	[classType] INTEGER NOT NULL DEFAULT 1,
	[o2tConnectionPoint] INT NOT NULL DEFAULT 1,
	[o2tSize] INT NOT NULL DEFAULT 496,
	[o2tRpi] INT NOT NULL DEFAULT 100,
	[o2tRtHeader] BOOLEAN NOT NULL DEFAULT 1,
	[t2oConnectionPoint] INT NOT NULL DEFAULT 2,
	[t2oSize] INT NOT NULL DEFAULT 496,
	[t2oRpi] INT NOT NULL DEFAULT 100,
	[t2oRtHeader] BOOLEAN NOT NULL DEFAULT 0,
	[t2oPriority] INT NOT NULL DEFAULT 2,
	[t2oConnectionType] BOOLEAN NOT NULL DEFAULT 0,
	[timeoutMultiplier] INT NOT NULL DEFAULT 2,
	[configurationInstance] INT NOT NULL DEFAULT 1, 
	[fpFunc] INT NOT NULL DEFAULT 0,
	[fpTimeout] INT NOT NULL DEFAULT 60000,
	[fpData] TEXT NOT NULL DEFAULT "00 00",
	CONSTRAINT [fkDevice] FOREIGN KEY ([deviceId]) REFERENCES [devices]([deviceId]) ON DELETE CASCADE ON UPDATE CASCADE,
	UNIQUE (deviceId, name),
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1)),
	CONSTRAINT [classType (1)] CHECK ([classType] IN (1)),
	CONSTRAINT [o2tConnectionPoint (1-2147483647)] CHECK ([o2tConnectionPoint] BETWEEN 1 AND 2147483647),
	CONSTRAINT [o2tSize (0-496)] CHECK ([o2tSize] BETWEEN 0 AND 496),
	CONSTRAINT [o2tRpi (1-3000)] CHECK ([o2tRpi] BETWEEN 1 AND 3000),
	CONSTRAINT [o2tRtHeader (0,1)] CHECK ([o2tRtHeader] IN (0,1)), /* 0 = Modeless, 1 = 32-Bit Header */
	CONSTRAINT [t2oConnectionPoint (1-2147483647)] CHECK ([t2oConnectionPoint] BETWEEN 1 AND 2147483647),
	CONSTRAINT [t2oSize (0-496)] CHECK ([t2oSize] BETWEEN 0 AND 496),
	CONSTRAINT [o2tSize >= 1 or t2oSize >= 1] CHECK ([o2tSize] >= 1 OR [t2oSize] >= 1),
	CONSTRAINT [t2oRpi (1-3000)] CHECK ([t2oRpi] BETWEEN 1 AND 3000),
	CONSTRAINT [t2oRtHeader (0,1)] CHECK ([t2oRtHeader] IN (0,1)), /* 0 = Modeless, 1 = 32-Bit Header */
	CONSTRAINT [t2oPriority (2)] CHECK ([t2oPriority] IN (2)),
	CONSTRAINT [t2oConnectionType (0)] CHECK ([t2oConnectionType] IN (0)), /* 0 = Point to Point, 1 = Multicast */
	CONSTRAINT [timeoutMultiplier (0-7)] CHECK ([timeoutMultiplier] BETWEEN 0 AND 7), /* 0=x4/1=x8/2=x16/3=x32/4=x64/5=x128/6=x256/7=x512 */
	CONSTRAINT [configurationInstance (1-2147483647)] CHECK ([configurationInstance] BETWEEN 1 AND 2147483647),
	CONSTRAINT [fpFunc (0,1,3)] CHECK ([fpFunc] IN (0,1,3)),
	CONSTRAINT [fpTimeout (100-65535)] CHECK ([fpTimeout] BETWEEN 100 AND 65535),
	CONSTRAINT [fpData length (0-1488)] CHECK (length(fpData) BETWEEN 0 AND 1488) /* 496x3 */
);
CREATE TABLE [devices](
	[deviceId] INTEGER PRIMARY KEY, 
	[scannerId] INTEGER NOT NULL,
	[name] TEXT NOT NULL DEFAULT "device", 
	[enable] BOOLEAN NOT NULL DEFAULT 1,
	[ipAddress] TEXT NOT NULL DEFAULT "0.0.0.0",
	[port] INT NOT NULL DEFAULT 44818,
	[ethernetId] INT NOT NULL DEFAULT 1,
	CONSTRAINT [fkScanner] FOREIGN KEY ([scannerId]) REFERENCES [scanners]([scannerId]) ON DELETE CASCADE ON UPDATE CASCADE,
	UNIQUE (name),
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1)),
	CONSTRAINT [ipAddress length (1-15)] CHECK (length(ipAddress) BETWEEN 1 AND 15),
	CONSTRAINT [port (1-65535)] CHECK ([port] BETWEEN 1 AND 65535),
	CONSTRAINT [ethernetId >= 1] CHECK ([ethernetId] >= 1)
);
CREATE TABLE [scanners](
	[scannerId] INTEGER PRIMARY KEY,
	[configId] INTEGER NOT NULL,
	[name] TEXT NOT NULL DEFAULT "eip_scanner",
	[enable] BOOLEAN NOT NULL DEFAULT 1,
	[enableDeviceFailEvent] BOOLEAN NOT NULL DEFAULT 1,
	[enableConnectionFailEvent] BOOLEAN NOT NULL DEFAULT 1,
	[eipEncapInactivityTimeoutValue] INT NOT NULL DEFAULT 120,
	CONSTRAINT [fkConfig] FOREIGN KEY ([configId]) REFERENCES [config]([configId]) ON DELETE CASCADE ON UPDATE CASCADE,
	UNIQUE (name),
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [enable (0,1)] CHECK ([enable] IN (0,1)),
	CONSTRAINT [enableDeviceFailEvent (0,1)] CHECK ([enableDeviceFailEvent] IN (0,1)),
	CONSTRAINT [enableConnectionFailEvent (0,1)] CHECK ([enableConnectionFailEvent] IN (0,1)),
	CONSTRAINT [eipEncapInactivityTimeoutValue (0-3600)] CHECK ([eipEncapInactivityTimeoutValue] BETWEEN 0 AND 3600)
);
CREATE TABLE [tags](
	[tagId] INTEGER PRIMARY KEY,
	[name] TEXT NOT NULL DEFAULT 'tag',
	[dataType] TEXT NOT NULL DEFAULT 'raw',	
	[dataUnit] TEXT NULL,
	[access] TEXT NOT NULL,
	[dataSize] INT NULL,
	/* start EIP params: */
	[connectionId] INTEGER NULL,
	[adapterId] INTEGER NULL,
	[byteOffset] INT NOT NULL DEFAULT 0,
	[bitOffset] INT NOT NULL DEFAULT 0,
	[scalingFunc] INT NOT NULL DEFAULT 0,
	[interceptSlope] REAL NOT NULL DEFAULT 1,
	[interceptOffset] REAL NOT NULL DEFAULT 0,
	[pointSourceMin] REAL NOT NULL DEFAULT 0,
	[pointSourceMax] REAL NOT NULL DEFAULT 1,
	[pointTargetMin] REAL NOT NULL DEFAULT 0,
	[pointTargetMax] REAL NOT NULL DEFAULT 1,
	CONSTRAINT [fkConnections] FOREIGN KEY ([connectionId]) REFERENCES [connections]([connectionId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [fkAdapter] FOREIGN KEY ([adapterId]) REFERENCES [adapters]([adapterId]) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT [name length (1-128)] CHECK (length(name) BETWEEN 1 AND 128),
	CONSTRAINT [name only accept - . _ ~ 0-9 a-z A-Z] CHECK ([name] NOT GLOB '*[^-0-9a-zA-Z._~]*'),
	CONSTRAINT [name should not be (status)] CHECK ([name] != 'status'),
	CONSTRAINT [dataType (boolean, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float, double, raw)] CHECK ([dataType] IN ("boolean", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float", "double", "raw")),	
	CONSTRAINT [dataUnit length <= 128] CHECK (length(dataUnit) <= 128),
	CONSTRAINT [access (r,w,rw)] CHECK ([access] IN ("r","w","rw")),
	CONSTRAINT [dataSize should exist when dataType is (raw), otherwise it should not exist] CHECK (([dataType] NOT IN ("raw")) OR (([dataType] IN ("raw")) AND ([dataSize] IS NOT NULL) AND ([dataSize] >= 1))),
	CONSTRAINT [bitOffset (0-7) when dataType is boolean, otherwise bitOffset (0)] CHECK (([dataType] NOT IN ("boolean") AND ([bitOffset] == 0)) OR ([dataType] IN ("boolean") AND ([bitOffset] BETWEEN 0 AND 7))),
	CONSTRAINT [scalingFunc (0,1,2)] CHECK ([scalingFunc] IN (0,1,2)),
	CONSTRAINT [exactly one of connectionId or adapterId] CHECK ((([connectionId] IS NULL) OR ([adapterId] IS NULL)) AND (([connectionId] IS NOT NULL) OR ([adapterId] IS NOT NULL)) )
);
 
-- INDEX
 
-- TRIGGER
CREATE TRIGGER [chkMaxAdaptersCount] 
BEFORE INSERT ON [adapters]
WHEN (
    SELECT COUNT(*) FROM [adapters] 
    ) >= 8
BEGIN
    SELECT RAISE(ABORT, "adapters are full(8)");
END;
CREATE TRIGGER [chkMaxConfigCount] 
BEFORE INSERT ON [config]
WHEN (
    SELECT COUNT(*) FROM [config] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "config are full(1)");
END;
CREATE TRIGGER [chkMaxConnectionsCount] 
BEFORE INSERT ON [connections]
WHEN (
    SELECT COUNT(*) FROM [connections] 
    ) >= 32
BEGIN
    SELECT RAISE(ABORT, "connections are full(32)");
END;
CREATE TRIGGER [chkMaxDevicesCount] 
BEFORE INSERT ON [devices]
WHEN (
    SELECT COUNT(*) FROM [devices] 
    ) >= 32
BEGIN
    SELECT RAISE(ABORT, "devices are full(32)");
END;
CREATE TRIGGER [chkMaxScannersCount] 
BEFORE INSERT ON [scanners]
WHEN (
    SELECT COUNT(*) FROM [scanners] 
    ) >= 1
BEGIN
    SELECT RAISE(ABORT, "scanners are full(1)");
END;
CREATE TRIGGER [chkMaxTagsCount] 
BEFORE INSERT ON [tags]
WHEN (
    SELECT COUNT(*) FROM [tags] 
    ) >= 512
BEGIN
    SELECT RAISE(ABORT, "tags are full(512)");
END;


CREATE TRIGGER chkTopicInsert
BEFORE INSERT ON [tags]
WHEN 
	(
		(
			(NEW.[connectionId] IS NOT NULL) 
			AND
			(
				SELECT COUNT(*)
				FROM (
					SELECT [connections].deviceId AS refDevId FROM [tags]
					INNER JOIN [connections] ON [tags].name == NEW.[name] AND [connections].connectionId == [tags].connectionId 
					) AS filterTable
				WHERE (SELECT [connections].deviceId FROM [connections] WHERE [connections].connectionId == NEW.connectionId AND [connections].deviceId == filterTable.refDevId)
			)
		)
		OR
		(
			(NEW.[adapterId] IS NOT NULL)
			AND
			(
				SELECT COUNT(*) FROM [tags]
				INNER JOIN [adapters] ON [adapters].adapterId == [tags].adapterId AND [tags].name == NEW.[name]				
			)
		)
	) > 0
BEGIN
    SELECT RAISE(ABORT, 'The set of following column(s) should be unique: deviceId, tags.name');
END;

CREATE TRIGGER chkTopicUpdate
BEFORE UPDATE ON [tags]
WHEN 
	(
		(
			(NEW.[connectionId] IS NOT NULL) 
			AND
			(
				SELECT COUNT(*)
				FROM (
					SELECT [connections].deviceId AS refDevId FROM [tags]
					INNER JOIN [connections] ON [tags].name == NEW.[name] AND OLD.[name] != NEW.[name] AND [connections].connectionId == [tags].connectionId 
					) AS filterTable
				WHERE (SELECT [connections].deviceId FROM [connections] WHERE [connections].connectionId == NEW.connectionId AND [connections].deviceId == filterTable.refDevId)
			)
		)
		OR
		(
			(NEW.[adapterId] IS NOT NULL)
			AND
			(
				SELECT COUNT(*) FROM [tags]
				INNER JOIN [adapters] ON [adapters].adapterId == [tags].adapterId AND [tags].name == NEW.[name]				
			)
		)
	) > 0
BEGIN
    SELECT RAISE(ABORT, 'The set of following column(s) should be unique: deviceId, tags.name');
END;
 
CREATE TRIGGER setDataSizeNullInsert
AFTER INSERT ON [tags]
WHEN (NEW.dataType NOT IN ('raw'))
BEGIN
   UPDATE tags SET dataSize = null WHERE ROWID = NEW.ROWID;
END;

CREATE TRIGGER setDataSizeNullUpdate
AFTER UPDATE ON [tags]
WHEN (NEW.dataType NOT IN ('raw'))
BEGIN
   UPDATE tags SET dataSize = null WHERE ROWID = NEW.ROWID;
END;

-- VIEW
 
