import React from "react";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import Grid from "@material-ui/core/Grid";
import { Input, InputLabel, FormLabel, TextField, Link } from "@material-ui/core";





const Login = () => {

  return (
    
          <Box padding={0} margin={0}>
              <Grid container style={{ backgroundColor: "#EEEAEA" }} spacing={0}>
                  <Grid item xs={6}>
            
                  <Container style={{ height: "100vh", backgroundColor: "#EEEAEA",   margin: "150px 0px 0px 50px",}}>
                      <Typography component="h1" variant="h4" style={{fontFamily: "poppins", fontSize: 40, fontWeight: "bold"}}>
            Sign in to Boardly
          </Typography>

                            <TextField
              variant="outlined"
              margin="normal"
              required
              style={{ width: "60vh", }} 
              id="email"
              label="Company's Email"
              name="email"
              autoComplete="email"
              autoFocus
            />
            <TextField
              variant="outlined"
              margin="normal"
              required
              style={{ width: "60vh", }} 
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
            />

                        
<Typography>
  <a href="#"> Forgot password?</a>
   
    </Typography>
    <br /> <br />
  

                          <Button variant="contained" style ={{backgroundColor: "#22ad80", color: "#ffffff", width: "60vh"}}
              type="submit">
              Sign In
            </Button> 



            
<Typography> New here? 
 <a href="/create-admin">  Create an account </a>
  </Typography>

</Container>
                     
                  </Grid>
                  <Grid item xs={6}>
                      <Box>
                      <Container style={{ height: "100vh", backgroundColor: "#753FF6",  clipPath: "polygon(8% 0, 100% 0, 100% 100%, 0% 100%)"  }}>
                          <div style={{
                                margin: "0 50px 0 50px",
                                position: 'absolute', left: '50%', top: '50%',
                                transform: 'translate(10%, -50%)'
                            }} >
                              <img src="/window.png"/>
  
                          </div>
                      </Container>
                      </Box>
                  </Grid>
              </Grid>
  
          </Box>
  );
};
export default Login;


