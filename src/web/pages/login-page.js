import React from "react";
import Container from "@material-ui/core/Container";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";
import Button from "@material-ui/core/Button";
import { makeStyles } from "@material-ui/core/styles";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import { FormControl, withStyles } from "@material-ui/core";
import { Input, InputLabel, FormLabel, TextField } from "@material-ui/core";
import InputBase from "@material-ui/core/InputBase";

const BootstrapInput = withStyles((theme) => ({
  root: {
    "label + &": {
      marginTop: theme.spacing(3),
    },
  },
  input: {
    borderRadius: 4,
    border: "1px solid black",
    position: "relative",
    // backgroundColor: theme.palette.common.white,
    border: "2px solid black",
    fontSize: 16,
    width: "auto",
    backgroundColor: "rgba(196, 196, 196, 0.2)",
    padding: "10px 12px",
    transition: theme.transitions.create(["border-color", "box-shadow"]),
    "&:focus": {
      // boxShadow: `${fade(theme.palette.primary.main, 0.25)} 0 0 0 0.2rem`,
      borderColor: theme.palette.primary.main,
    },
  },
}))(InputBase);

const useStyles = makeStyles((theme) => ({
  root: {
    // display: "flex",
    // flexWrap: "wrap",
  },
  margin: {
    margin: theme.spacing(1),
  },
}));

const Login = () => {
  const classes = useStyles();
  return (
    <div>
      {/* <Grid item xs={6}>
          <Box>
            <Typography>Sign in to Boardly</Typography>
            <FormControl>
              <FormLabel>Company's Name</FormLabel>
              <Input id="my-input" placeholder="user@boardly.com" />
            </FormControl>{" "}
            <br />
            <FormControl>
              <FormLabel>Password</FormLabel>
              <Input id="my-input" placeholder="Password" />
            </FormControl>
            <TextField
              id="username"
              type="email"
              label="Company's Email"
              placeholder="Username"
              margin="normal"
            />
          </Box>
        </Grid>
        <Grid item xs={6}>
          <Paper className={classes.paper}>xs=6</Paper>
        </Grid> */}

      <Grid
        container
        direction="row"
        alignItems="flex-start"
        classes={{ gridTemplateColumns: "50% 50%" }}
      >
        <Box color="text.primary" m={1}>
          <img src="./boardly_logo.png" />
          <Typography variant="h3" gutterBottom>
            Sign In to Boardly
          </Typography>

          <form noValidate>
            <FormControl>
              <InputLabel htmlFor="bootstrap-input">Company's Email</InputLabel>
              <BootstrapInput
                defaultValue="react-bootstrap"
                id="bootstrap-input"
              />
            </FormControl>

            <FormControl>
              <InputLabel htmlFor="bootstrap-input">Password</InputLabel>
              <BootstrapInput
                defaultValue="react-bootstrap"
                id="bootstrap-input"
              />
            </FormControl>

            <Typography variant="h6">Forget Password ?</Typography>
          </form>
        </Box>

        <Box m={1}>
          <Button variant="contained" color="primary">
            Hello World
          </Button>
          
          <Button>Click me too! </Button>
        </Box>
      </Grid>
    </div>
  );
};
export default Login;
