// Code generated by sql2go. DO NOT EDIT.
// source: schema.sql

package storage

var sqlSchemaCmds = [...]string{
	"CREATE TABLE IF NOT EXISTS keyvals ( key TEXT PRIMARY KEY NOT NULL, tbl TEXT NOT NULL, val BYTEA NOT NULL );",
	"CREATE INDEX IF NOT EXISTS kvs_tbl_idx ON keyvals (tbl);",
	"CREATE TABLE IF NOT EXISTS channels ( cid TEXT PRIMARY KEY NOT NULL, peer TEXT NOT NULL, token TEXT NOT NULL, ledger TEXT NOT NULL, state INT NOT NULL, statets TIMESTAMPTZ NOT NULL, opents TIMESTAMPTZ NOT NULL, openresp BYTEA, onchainbalance BYTEA, basesn INT NOT NULL, lastusedsn INT NOT NULL, lastackedsn INT NOT NULL, lastnackedsn INT NOT NULL, selfsimplex BYTEA, peersimplex BYTEA, UNIQUE (peer, token) );",
	"CREATE INDEX IF NOT EXISTS chan_ledger_idx ON channels (ledger);",
	"CREATE INDEX IF NOT EXISTS chan_state_idx ON channels (state);",
	"CREATE INDEX IF NOT EXISTS chan_token_idx ON channels (token);",
	"CREATE TABLE IF NOT EXISTS closedchannels ( cid TEXT PRIMARY KEY NOT NULL, peer TEXT NOT NULL, token TEXT NOT NULL, opents TIMESTAMPTZ NOT NULL, closets TIMESTAMPTZ NOT NULL );",
	"CREATE INDEX IF NOT EXISTS cc_peer_token_idx ON closedchannels (peer, token);",
	"CREATE TABLE IF NOT EXISTS payments ( payid TEXT PRIMARY KEY NOT NULL, pay BYTEA, paynote BYTEA, incid TEXT NOT NULL, instate INT NOT NULL, outcid TEXT NOT NULL, outstate INT NOT NULL, src TEXT NOT NULL, dest TEXT NOT NULL, createts TIMESTAMPTZ NOT NULL );",
	"CREATE INDEX IF NOT EXISTS pay_src_idx ON payments (src);",
	"CREATE INDEX IF NOT EXISTS pay_dest_idx ON payments (dest);",
	"CREATE INDEX IF NOT EXISTS pay_ts_idx ON payments (createts);",
	"CREATE TABLE IF NOT EXISTS paydelegation ( payid TEXT PRIMARY KEY NOT NULL REFERENCES payments (payid) ON UPDATE CASCADE ON DELETE CASCADE, dest TEXT NOT NULL, status INT NOT NULL, payidout TEXT, delegator TEXT );",
	"CREATE INDEX IF NOT EXISTS paydel_dest_idx ON paydelegation (dest);",
	"CREATE TABLE IF NOT EXISTS crossnetpays ( payid TEXT PRIMARY KEY NOT NULL REFERENCES payments (payid) ON UPDATE CASCADE ON DELETE CASCADE, originalpayid TEXT NOT NULL, originalpay BYTEA, state INT NOT NULL, srcnetid INT NOT NULL, dstnetid INT NOT NULL, bridgeaddr TEXT NOT NULL, bridgenetid INT NOT NULL, UNIQUE (originalpayid) );",
	"CREATE TABLE IF NOT EXISTS secrets ( hash TEXT PRIMARY KEY NOT NULL, preimage TEXT NOT NULL, payid TEXT NOT NULL, UNIQUE (hash, payid) );",
	"CREATE TABLE IF NOT EXISTS tcb ( addr TEXT NOT NULL, token TEXT NOT NULL, deposit TEXT NOT NULL, UNIQUE (addr, token) );",
	"CREATE TABLE IF NOT EXISTS monitor ( event TEXT PRIMARY KEY NOT NULL, blocknum INT NOT NULL, blockidx INT NOT NULL, restart BOOL NOT NULL );",
	"CREATE TABLE IF NOT EXISTS routing ( dest TEXT NOT NULL, token TEXT NOT NULL, cid TEXT NOT NULL, UNIQUE (dest, token) );",
	"CREATE TABLE IF NOT EXISTS edges ( cid TEXT PRIMARY KEY NOT NULL, token TEXT NOT NULL, addr1 TEXT NOT NULL, addr2 TEXT NOT NULL );",
	"CREATE INDEX IF NOT EXISTS edges_token_idx ON edges (token);",
	"CREATE TABLE IF NOT EXISTS netbridge ( bridgeaddr TEXT PRIMARY KEY NOT NULL, bridgenetid INT NOT NULL );",
	"CREATE TABLE IF NOT EXISTS bridgerouting ( destnetid INT PRIMARY KEY NOT NULL, bridgeaddr TEXT NOT NULL );",
	"CREATE TABLE IF NOT EXISTS nettokens ( netid INT NOT NULL, nettoken TEXT NOT NULL, localtoken TEXT NOT NULL, rate FLOAT NOT NULL, UNIQUE (netid, nettoken), UNIQUE (netid, localtoken) );",
	"CREATE TABLE IF NOT EXISTS peers ( peer TEXT PRIMARY KEY NOT NULL, server TEXT NOT NULL, activecids TEXT NOT NULL,  delegateproof BYTEA );",
	"CREATE INDEX IF NOT EXISTS peers_server_idx ON peers (server);",
	"CREATE TABLE IF NOT EXISTS desttokens ( dest TEXT NOT NULL, token TEXT NOT NULL, osps TEXT NOT NULL,  openchanblknum INT NOT NULL, UNIQUE (dest, token) );",
	"CREATE TABLE IF NOT EXISTS chanmessages ( cid TEXT NOT NULL, seqnum INT NOT NULL, msg BYTEA, UNIQUE (cid, seqnum) );",
	"CREATE TABLE IF NOT EXISTS chanmigration ( cid TEXT NOT NULL REFERENCES channels (cid) ON DELETE CASCADE, toledger TEXT NOT NULL, deadline INT NOT NULL, onchainreq BYTEA, state INT NOT NULL, ts TIMESTAMPTZ NOT NULL, UNIQUE (cid, toledger) );",
	"CREATE INDEX IF NOT EXISTS mg_toledger_state_idx ON chanmigration (toledger, state);",
	"CREATE TABLE IF NOT EXISTS deposit ( uuid TEXT PRIMARY KEY NOT NULL, cid TEXT NOT NULL, topeer BOOL NOT NULL, amount TEXT NOT NULL, refill BOOL NOT NULL, deadline TIMESTAMPTZ NOT NULL, state INT NOT NULL, txhash TEXT NOT NULL, errmsg TEXT NOT NULL );",
	"CREATE INDEX IF NOT EXISTS deposit_cid_idx ON deposit (cid);",
	"CREATE INDEX IF NOT EXISTS deposit_state_idx ON deposit (state);",
	"CREATE INDEX IF NOT EXISTS deposit_txhash_idx ON deposit (txhash);",
	"CREATE TABLE IF NOT EXISTS lease ( id TEXT PRIMARY KEY NOT NULL, owner TEXT NOT NULL, updatets TIMESTAMPTZ NOT NULL );",
}
