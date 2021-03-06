import React from 'react';
// import { Button } from '@material-ui/core';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { 
    FormControl, 
    Grid, 
    Typography, 
    Box,
    Input 
} from '@material-ui/core';
import RightBackground from "../../public/Group5.svg" 
import Forgot_password from "../../public/forgot_password.svg" 
import Logo from "../../public/logo.svg"
import Button from '@material-ui/core/Button';
import { green } from '@material-ui/core/colors';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
      position: "fixed",
      height: "100%",
      width:"100%",
      margin :"0px",
      padding:"0px",
      top:"0px",
      backgroundColor: "#E5E5E5",
    },
    
    rightGrid: {
        position: "relative",
        backgroundColor: "#E5E5E5",
        height: "100%",
    },
    leftGrid: {
        height: "100%",
        backgroundColor: "#E5E5E5",
        // border:"1px solid red",
        padding: 0,
        margin:0,
    },
    forgotPassHeader: {
        fontFamily: "Poppins",
        fontStyle: "normal",
        fontWeight: "bold",
        fontSize: "40px",
        lineHeight: "60px",
        marginLeft:"10%"
    },
    forgotPassInput: {

    }

  }),
);

export const ForgotPassword: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            
        <div className={classes.root}>
            <Grid container spacing={3}>
                <Grid item xs={12} sm={6} className={classes.leftGrid}>
                        <img  src={Logo} width="50px" height="50px"
                        style={{
                            marginLeft:"10%",
                            marginTop:"10%",
                        }}
                        />
                        
                        <Box className={classes.forgotPassHeader}>
                               <p>Forgot Password</p>
                        </Box>

                        <Box className={classes.forgotPassInput} m="10%" mt="20%">
                             <Typography variant="h6" gutterBottom>
                                <p>E-mail address</p>
                            </Typography>

                            <form noValidate autoComplete="off">
                                <FormControl fullWidth>
                                    <Input placeholder="user@email.com" fullWidth
                                     style={{
                                        border:"0.2px solid black",
                                        width:"100%",
                                        borderRadius:"10px",
                                        backgroundColor: "#C4C4C4",
                                        padding:"10px"
                                    }} 
                                    />
                                </FormControl>
                            </form>

                            <Box mt="10%">
                                <Box display="flex" justifyContent="center" m={1} p={1} bgcolor="#E5E5E5">
                                        <Button variant="outlined" 
                                            style= {{
                                                width:"100%",
                                                backgroundColor:"#22AD80",
                                                color:"#ffff",
                                                fontSize:"12px",
                                                fontWeight:"bold",
                                                padding:"12px",
                                                borderRadius:"10px",

                                            }} >
                                            Send Reset Link
                                        </Button>
                                </Box>
                            </Box>

                        </Box>

                       


                  
                       
                        
                </Grid>
                <Grid item xs={12} sm={6} className={classes.rightGrid}>
                    <img style={{
                        height:"auto",
                        maxWidth:"100%",
                        position: "absolute",
                        top: "25%",
                        left: "25%",
                        }} 
                        src={Forgot_password} 
                    />
                    <img style={{
                        height:"auto",
                        width:"100%",
                        backgroundColor: "#E5E5E5",
                        border:"0"
                            }} 
                        src={RightBackground} 
                    />

                </Grid>
            </Grid>
        </div>

        </>
    );
};
