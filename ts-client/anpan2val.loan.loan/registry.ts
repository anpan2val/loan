import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { MsgRequestLoan } from "./types/loan/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/anpan2val.loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/anpan2val.loan.loan.MsgRequestLoan", MsgRequestLoan],
    
];

export { msgTypes }