import React from 'react';
import { ThemeProvider as MuiThemeProvider } from '@material-ui/core/styles';
import theme from '../theme';
import '../styles/globals.css';
import { AppProps } from 'next/dist/next-server/lib/router/router';

function MyApp({ Component, pageProps }: AppProps): JSX.Element {
    return (
        <MuiThemeProvider theme={theme}>
            <Component {...pageProps} />
        </MuiThemeProvider>
    );
}

export default MyApp;
