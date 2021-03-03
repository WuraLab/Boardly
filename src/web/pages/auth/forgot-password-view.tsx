import {ForgotPassword} from "../../components/forgotpassword/ForgotPassword"
import {Layout} from "../../components/layout/index"
import React from "react"


// export const ForgotPasswordView: React.FC = () => {
    // export default function ForgotPasswordView() {

const ForgotPasswordView: React.FC  = () => {

    return(
       
        <>
            <Layout children= {ForgotPassword}  />
        </>
    )


};

export default ForgotPasswordView