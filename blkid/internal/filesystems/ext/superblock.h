#include <stdint.h>

struct ext2_super_block {
	uint32_t		s_inodes_count;
	uint32_t		s_blocks_count;
	uint32_t		s_r_blocks_count;
	uint32_t		s_free_blocks_count;
	uint32_t		s_free_inodes_count;
	uint32_t		s_first_data_block;
	uint32_t		s_log_block_size;
	uint32_t		s_dummy3[7];
	unsigned char		s_magic[2];
	uint16_t		s_state;
	uint16_t		s_errors;
	uint16_t		s_minor_rev_level;
	uint32_t		s_lastcheck;
	uint32_t		s_checkinterval;
	uint32_t		s_creator_os;
	uint32_t		s_rev_level;
	uint16_t		s_def_resuid;
	uint16_t		s_def_resgid;
	uint32_t		s_first_ino;
	uint16_t		s_inode_size;
	uint16_t		s_block_group_nr;
	uint32_t		s_feature_compat;
	uint32_t		s_feature_incompat;
	uint32_t		s_feature_ro_compat;
	unsigned char		s_uuid[16];
	char			s_volume_name[16];
	char			s_last_mounted[64];
	uint32_t		s_algorithm_usage_bitmap;
	uint8_t			s_prealloc_blocks;
	uint8_t			s_prealloc_dir_blocks;
	uint16_t		s_reserved_gdt_blocks;
	uint8_t			s_journal_uuid[16];
	uint32_t		s_journal_inum;
	uint32_t		s_journal_dev;
	uint32_t		s_last_orphan;
	uint32_t		s_hash_seed[4];
	uint8_t			s_def_hash_version;
	uint8_t			s_jnl_backup_type;
	uint16_t		s_reserved_word_pad;
	uint32_t		s_default_mount_opts;
	uint32_t		s_first_meta_bg;
	uint32_t		s_mkfs_time;
	uint32_t		s_jnl_blocks[17];
	uint32_t		s_blocks_count_hi;
	uint32_t		s_r_blocks_count_hi;
	uint32_t		s_free_blocks_hi;
	uint16_t		s_min_extra_isize;
	uint16_t		s_want_extra_isize;
	uint32_t		s_flags;
	uint16_t		s_raid_stride;
	uint16_t		s_mmp_interval;
	uint64_t		s_mmp_block;
	uint32_t		s_raid_stripe_width;
	uint32_t		s_reserved[162];
	uint32_t		s_checksum;
} __attribute__((packed));
