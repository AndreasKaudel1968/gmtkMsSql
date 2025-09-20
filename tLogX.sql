use [DW]
go

/****** Object:  Table [dbo].[tLogX]    Script Date: 20.09.2025 20:17:28 ******/
set ansi_nulls on
go

set quoted_identifier on
go

create table [dbo].[tLogX](
	[id] [int] identity(1,1) not null,
	[src] [varchar](100) not null,
	[lvl] [tinyint] null,
	[task] [varchar](50) null,
	[obj] [varchar](250) null,
	[evtID] [int] null,
	[msg] [nvarchar](max) null,
	[createdAt] [datetime] not null,
	[createdBy] [varchar](50) null,
	[closedAt] [datetime] null,
	[duration] [time](3) null,
	[outdated] [datetime] null,
 constraint [PK_tLogX] primary key nonclustered 
(
	[id] asc
)with (pad_index = off, statistics_norecompute = off, ignore_dup_key = off, allow_row_locks = on, allow_page_locks = on, fillfactor = 95, optimize_for_sequential_key = off) on [PRIMARY]
) on [PRIMARY] textimage_on [PRIMARY]
go


