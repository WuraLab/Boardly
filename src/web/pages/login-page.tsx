import React from "react";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import { FormControl, withStyles } from "@material-ui/core";
import { Input, InputLabel, FormLabel, TextField, Link } from "@material-ui/core";
import InputBase from "@material-ui/core/InputBase";




const Login = () => {

  return (
    
          <Box padding={0} margin={0}>
              <Grid container style={{ backgroundColor: "white" }} spacing={0}>
                  <Grid item xs={6}>
            
                  <Container style={{ height: "100vh", backgroundColor: "#EEEAEA" }}>
                      <Typography component="h1" variant="h4" style={{fontFamily: "poppins", fontSize: 40, fontWeight: "bold"}}>
            Sign in to Boardly
          </Typography>

                            <TextField
              variant="outlined"
              margin="normal"
              required
              fullWidth
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
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
            />

                          <Grid container>
<Grid item xs>
  <Link href="#" variant="body2" style={{color:"#577B70"}}>
    Forgot password?
  </Link>
</Grid>
</Grid>
                          <Button style ={{color: "white",  fontFamily: "poppins", backgroundColor: "#22AD80"}}
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
             
            >
              Sign In
            </Button>

            <Grid container>
<Grid item > New here?  
   <Link href="#" variant="body2" style={{color:"#577B70"}}> 
       Create an account
  </Link>
</Grid>
</Grid>
                      </Container>
                  </Grid>
                  <Grid item xs={6}>
                      <Box>
                      <Container className="slanted" style={{ height: "100vh", backgroundColor: "#753FF6",  }}>
                          <Box justifyContent="center">
                              <img src="/window.png"/>
  
                          </Box>
                      </Container>
                      </Box>
                  </Grid>
              </Grid>
  
          </Box>
  );
};
export default Login;


