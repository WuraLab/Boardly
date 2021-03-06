// 1. GlobalStyles.js
import { createStyles, makeStyles } from '@material-ui/core';
import React from 'react';

const useStyles = makeStyles(() =>
  createStyles({
    '@global': {
      html: {
        '-webkit-font-smoothing': 'antialiased',
        '-moz-osx-font-smoothing': 'grayscale',
        height: '100%',
        width: '100%'
      },
      '*, *::before, *::after': {
        boxSizing: 'inherit'
      },
      body: {
        height: '100%',
        width: '100%',
        margin :"0px",
        background: "E5E5E5",
        fontFamily: "Poppins"
      },
    //   '#root': {
    //     height: '100%',
    //     width: '100%'
    //   }
    }
  })
);

const GlobalStyles: React.FC = () => {
  useStyles();

  return null;
};

export default GlobalStyles;