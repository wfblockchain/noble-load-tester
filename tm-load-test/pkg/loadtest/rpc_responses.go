package loadtest

import "time"

type TxResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Query string `json:"query"`
		Data  struct {
			Type  string `json:"type"`
			Value struct {
				TxResult struct {
					Height string `json:"height"`
					Tx     string `json:"tx"`
					Result struct {
						Data      string `json:"data"`
						Log       string `json:"log"`
						GasWanted string `json:"gas_wanted"`
						GasUsed   string `json:"gas_used"`
						Events    []struct {
							Type       string `json:"type"`
							Attributes []struct {
								Key   string `json:"key"`
								Value string `json:"value"`
								Index bool   `json:"index"`
							} `json:"attributes"`
						} `json:"events"`
					} `json:"result"`
				} `json:"TxResult"`
			} `json:"value"`
		} `json:"data"`
		Events struct {
			TxFee                []string `json:"tx.fee"`
			TxFeePayer           []string `json:"tx.fee_payer"`
			CoinSpentSpender     []string `json:"coin_spent.spender"`
			CoinReceivedAmount   []string `json:"coin_received.amount"`
			TransferSender       []string `json:"transfer.sender"`
			TxHash               []string `json:"tx.hash"`
			TxSignature          []string `json:"tx.signature"`
			CoinSpentAmount      []string `json:"coin_spent.amount"`
			TransferAmount       []string `json:"transfer.amount"`
			CoinReceivedReceiver []string `json:"coin_received.receiver"`
			TransferRecipient    []string `json:"transfer.recipient"`
			TmEvent              []string `json:"tm.event"`
			TxHeight             []string `json:"tx.height"`
			TxAccSeq             []string `json:"tx.acc_seq"`
			MessageAction        []string `json:"message.action"`
			MessageSender        []string `json:"message.sender"`
			MessageModule        []string `json:"message.module"`
		} `json:"events"`
	} `json:"result"`
}

type Block struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Query string `json:"query"`
		Data  struct {
			Type  string `json:"type"`
			Value struct {
				Block struct {
					Header struct {
						Version struct {
							Block string `json:"block"`
						} `json:"version"`
						ChainID     string    `json:"chain_id"`
						Height      string    `json:"height"`
						Time        time.Time `json:"time"`
						LastBlockID struct {
							Hash  string `json:"hash"`
							Parts struct {
								Total int    `json:"total"`
								Hash  string `json:"hash"`
							} `json:"parts"`
						} `json:"last_block_id"`
						LastCommitHash     string `json:"last_commit_hash"`
						DataHash           string `json:"data_hash"`
						ValidatorsHash     string `json:"validators_hash"`
						NextValidatorsHash string `json:"next_validators_hash"`
						ConsensusHash      string `json:"consensus_hash"`
						AppHash            string `json:"app_hash"`
						LastResultsHash    string `json:"last_results_hash"`
						EvidenceHash       string `json:"evidence_hash"`
						ProposerAddress    string `json:"proposer_address"`
					} `json:"header"`
					Data struct {
						Txs []string `json:"txs"`
					} `json:"data"`
					Evidence struct {
						Evidence []any `json:"evidence"`
					} `json:"evidence"`
					LastCommit struct {
						Height  string `json:"height"`
						Round   int    `json:"round"`
						BlockID struct {
							Hash  string `json:"hash"`
							Parts struct {
								Total int    `json:"total"`
								Hash  string `json:"hash"`
							} `json:"parts"`
						} `json:"block_id"`
						Signatures []struct {
							BlockIDFlag      int       `json:"block_id_flag"`
							ValidatorAddress string    `json:"validator_address"`
							Timestamp        time.Time `json:"timestamp"`
							Signature        string    `json:"signature"`
						} `json:"signatures"`
					} `json:"last_commit"`
				} `json:"block"`
				ResultBeginBlock struct {
					Events []struct {
						Type       string `json:"type"`
						Attributes []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
							Index bool   `json:"index"`
						} `json:"attributes"`
					} `json:"events"`
				} `json:"result_begin_block"`
				ResultEndBlock struct {
					ValidatorUpdates []struct {
						PubKey struct {
							Sum struct {
								Type  string `json:"type"`
								Value struct {
									Ed25519 string `json:"ed25519"`
								} `json:"value"`
							} `json:"Sum"`
						} `json:"pub_key"`
						Power string `json:"power"`
					} `json:"validator_updates"`
					ConsensusParamUpdates struct {
						Block struct {
							MaxBytes string `json:"max_bytes"`
							MaxGas   string `json:"max_gas"`
						} `json:"block"`
						Evidence struct {
							MaxAgeNumBlocks string `json:"max_age_num_blocks"`
							MaxAgeDuration  string `json:"max_age_duration"`
							MaxBytes        string `json:"max_bytes"`
						} `json:"evidence"`
						Validator struct {
							PubKeyTypes []string `json:"pub_key_types"`
						} `json:"validator"`
					} `json:"consensus_param_updates"`
					Events []struct {
						Type       string `json:"type"`
						Attributes []struct {
							Key   string `json:"key"`
							Value string `json:"value"`
							Index bool   `json:"index"`
						} `json:"attributes"`
					} `json:"events"`
				} `json:"result_end_block"`
			} `json:"value"`
		} `json:"data"`
		Events struct {
			SendPacketPacketConnection       []string `json:"send_packet.packet_connection"`
			TmEvent                          []string `json:"tm.event"`
			CoinReceivedAmount               []string `json:"coin_received.amount"`
			MessageSender                    []string `json:"message.sender"`
			MintBondedRatio                  []string `json:"mint.bonded_ratio"`
			SendPacketPacketDstPort          []string `json:"send_packet.packet_dst_port"`
			CoinSpentAmount                  []string `json:"coin_spent.amount"`
			MintInflation                    []string `json:"mint.inflation"`
			SendPacketPacketData             []string `json:"send_packet.packet_data"`
			RewardsValidator                 []string `json:"rewards.validator"`
			SendPacketPacketTimeoutHeight    []string `json:"send_packet.packet_timeout_height"`
			SendPacketPacketChannelOrdering  []string `json:"send_packet.packet_channel_ordering"`
			SendPacketConnectionID           []string `json:"send_packet.connection_id"`
			SendPacketPacketTimeoutTimestamp []string `json:"send_packet.packet_timeout_timestamp"`
			SendPacketPacketDstChannel       []string `json:"send_packet.packet_dst_channel"`
			TransferRecipient                []string `json:"transfer.recipient"`
			TransferAmount                   []string `json:"transfer.amount"`
			MintAnnualProvisions             []string `json:"mint.annual_provisions"`
			CommissionAmount                 []string `json:"commission.amount"`
			CoinbaseAmount                   []string `json:"coinbase.amount"`
			CoinSpentSpender                 []string `json:"coin_spent.spender"`
			MessageModule                    []string `json:"message.module"`
			MintAmount                       []string `json:"mint.amount"`
			CommissionValidator              []string `json:"commission.validator"`
			RewardsAmount                    []string `json:"rewards.amount"`
			SendPacketPacketDataHex          []string `json:"send_packet.packet_data_hex"`
			SendPacketPacketSrcPort          []string `json:"send_packet.packet_src_port"`
			SendPacketPacketSrcChannel       []string `json:"send_packet.packet_src_channel"`
			CoinReceivedReceiver             []string `json:"coin_received.receiver"`
			CoinbaseMinter                   []string `json:"coinbase.minter"`
			TransferSender                   []string `json:"transfer.sender"`
			SendPacketPacketSequence         []string `json:"send_packet.packet_sequence"`
		} `json:"events"`
	} `json:"result"`
}
