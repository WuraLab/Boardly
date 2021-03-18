import React from 'react';
import { ForgotPassword } from '../../components/forgotpassword/ForgotPassword';
// import { Layout } from '../../components/layout/index';
import GlobalStyles from '../../assets/global-styles';

// export const ForgotPasswordView: React.FC = () => {
// export default function ForgotPasswordView() {

const ForgotPasswordView: React.FC = () => {
    return (
        // <Layout>
        <>
            <GlobalStyles />
            <ForgotPassword />
            {/* // </Layout> */}
        </>
    );
};

export default ForgotPasswordView;
