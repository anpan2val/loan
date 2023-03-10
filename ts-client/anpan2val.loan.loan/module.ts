// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";


export { MsgApproveLoan, MsgRequestLoan };

type sendMsgApproveLoanParams = {
  value: MsgApproveLoan,
  fee?: StdFee,
  memo?: string
};

type sendMsgRequestLoanParams = {
  value: MsgRequestLoan,
  fee?: StdFee,
  memo?: string
};


type msgApproveLoanParams = {
  value: MsgApproveLoan,
};

type msgRequestLoanParams = {
  value: MsgRequestLoan,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgApproveLoan({ value, fee, memo }: sendMsgApproveLoanParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveLoan: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveLoan({ value: MsgApproveLoan.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveLoan: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRequestLoan({ value, fee, memo }: sendMsgRequestLoanParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRequestLoan: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRequestLoan({ value: MsgRequestLoan.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRequestLoan: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgApproveLoan({ value }: msgApproveLoanParams): EncodeObject {
			try {
				return { typeUrl: "/anpan2val.loan.loan.MsgApproveLoan", value: MsgApproveLoan.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveLoan: Could not create message: ' + e.message)
			}
		},
		
		msgRequestLoan({ value }: msgRequestLoanParams): EncodeObject {
			try {
				return { typeUrl: "/anpan2val.loan.loan.MsgRequestLoan", value: MsgRequestLoan.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRequestLoan: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			Anpan2ValLoanLoan: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;