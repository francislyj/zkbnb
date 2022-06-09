package svc

import (
	"github.com/zecrey-labs/zecrey-legend/common/model/account"
	"github.com/zecrey-labs/zecrey-legend/common/model/block"
	"github.com/zecrey-labs/zecrey-legend/common/model/l1TxSender"
	"github.com/zecrey-labs/zecrey-legend/common/model/liquidity"
	"github.com/zecrey-labs/zecrey-legend/common/model/nft"
	"github.com/zecrey-labs/zecrey-legend/common/model/proofSender"
	"github.com/zecrey-labs/zecrey-legend/common/model/sysconfig"
	"github.com/zecrey-labs/zecrey-legend/service/rpc/proverHub/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	BlockModel            block.BlockModel
	L1TxSenderModel       l1TxSender.L1TxSenderModel
	SysConfigModel        sysconfig.SysconfigModel
	AccountModel          account.AccountModel
	AccountHistoryModel   account.AccountHistoryModel
	LiquidityModel        liquidity.LiquidityModel
	LiquidityHistoryModel liquidity.LiquidityHistoryModel
	NftModel              nft.L2NftModel
	NftHistoryModel       nft.L2NftHistoryModel
	SysconfigModel        sysconfig.SysconfigModel
	ProofSenderModel      proofSender.ProofSenderModel
}

func WithRedis(redisType string, redisPass string) redis.Option {
	return func(p *redis.Redis) {
		p.Type = redisType
		p.Pass = redisPass
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	gormPointer, err := gorm.Open(postgres.Open(c.Postgres.DataSource))
	if err != nil {
		logx.Errorf("gorm connect db error, err = %s", err.Error())
	}
	conn := sqlx.NewSqlConn("postgres", c.Postgres.DataSource)
	redisConn := redis.New(c.CacheRedis[0].Host, WithRedis(c.CacheRedis[0].Type, c.CacheRedis[0].Pass))
	return &ServiceContext{
		Config:                c,
		BlockModel:            block.NewBlockModel(conn, c.CacheRedis, gormPointer, redisConn),
		L1TxSenderModel:       l1TxSender.NewL1TxSenderModel(conn, c.CacheRedis, gormPointer),
		SysConfigModel:        sysconfig.NewSysconfigModel(conn, c.CacheRedis, gormPointer),
		AccountModel:          account.NewAccountModel(conn, c.CacheRedis, gormPointer),
		AccountHistoryModel:   account.NewAccountHistoryModel(conn, c.CacheRedis, gormPointer),
		LiquidityModel:        liquidity.NewLiquidityModel(conn, c.CacheRedis, gormPointer),
		LiquidityHistoryModel: liquidity.NewLiquidityHistoryModel(conn, c.CacheRedis, gormPointer),
		NftModel:              nft.NewL2NftModel(conn, c.CacheRedis, gormPointer),
		NftHistoryModel:       nft.NewL2NftHistoryModel(conn, c.CacheRedis, gormPointer),
		SysconfigModel:        sysconfig.NewSysconfigModel(conn, c.CacheRedis, gormPointer),
		ProofSenderModel:      proofSender.NewProofSenderModel(gormPointer),
	}
}
