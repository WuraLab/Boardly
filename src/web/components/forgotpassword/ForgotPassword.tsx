import React from 'react';
// import { Button } from '@material-ui/core';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Container from '@material-ui/core/Container';
import { ReactSVG } from 'react-svg'
import Group5 from "../../public/Group5.svg" 
import Forgot_password from "../../public/forgot_password.svg" 

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1
    },
    paper: {
      padding: 0,
      margin:0,
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    rightGrid: {
        position: "relative",
        textAlign: "center",
        color: "white",
    }
  }),
);

export const ForgotPassword: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            {/* <p>Password Recovery component</p>
            <Button color="primary">Hello World</Button> */}
            {/* <Container maxWidth="sm"> */}
        <div className={classes.root}>
            <Grid container spacing={3}>
                <Grid item xs={12} sm={6}>
                    <Paper className={classes.paper}>
                        
                        xs=12 sm=6
                    
                    </Paper>
                </Grid>
                <Grid item xs={12} sm={6} className={classes.rightGrid}>
                    <Paper className={classes.paper}>
                    <img style={{
                        height:"auto",
                        maxWidth:"100%",
                        position: "absolute",
                        top: "50%",
                        left: "50%",
                        }} 
                        src={Forgot_password} 
                    />
                    <img style={{
                        height:"auto",
                        maxWidth:"100%"
                            }} 
                        src={Group5} 
                    />
                        
                    </Paper>
                </Grid>
            </Grid>
        </div>

        </>
    );
};
