import Box from "@material-ui/core/Box";
import Grid from "@material-ui/core/Grid";
import { Container, FormControl, Button, Typography } from "@material-ui/core";
import "@fontsource/poppins";
import InputLabel from "@material-ui/core/InputLabel";
import {
    createStyles,
    Theme,
    makeStyles,
} from "@material-ui/core/styles";
import { BootstrapInput } from "./BootstrapInput";
import { FormInput } from "./FormInput";

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: "flex",
            flexWrap: "wrap",
        },
        margin: {
            margin: theme.spacing(1),
        },
        input: {
            fontWeight: "bold",
            fontSize: "20px",
            fontFamily: "poppins",
            color: "black"
        }
    }),
);


const CreateAdmin = () => {
    const classes = useStyles();
    return (
        <Box padding={0} margin={0}>
            <Grid container spacing={0}>
                <Grid item xs={6}>
                    <Box style={{ margin: "20px 0 20px 100px" }}>
                        <img src="/logo.png" width="100" height="100" />
                    </Box>
                    <Container style={{ height: "100vh", backgroundColor: "white", display: "flex", margin: "50px 0 0 0" }} >

                        <Grid item container direction="column" alignItems="center">
                            <Typography style={{ color: "#000000", fontWeight: "bold", fontSize: "30px", letterSpacing: "4px" }}>
                                Create Admin Account
                            </Typography>
                            <Typography style={{ color: "#867979" }}>Fill out  the folowing information  to create an admin account</Typography>
                            <br />
                            <FormInput name="Company's Name" placeholder="Wuralab"> </FormInput>
                            <FormInput name="Company's Email" placeholder="contact@wuralab.com"> </FormInput>
                            <FormInput name="Password" placeholder="Password"> </FormInput>
                            <FormInput name="Confirm Password" placeholder="confirm-password"> </FormInput>
                            < br />
                            <Button variant="contained" style={{ backgroundColor: "#22ad80", color: "#ffffff", width: "50vh" }} >
                                Create Account
                            </Button>
                            <Typography color="primary">
                                <a href="/login"> Have an account?  <span> Login </span>  </a>
                            </Typography>
                        </Grid>

                    </Container>
                </Grid>
                <Grid item xs={6}>
                    <Box>
                        <Container style={{ height: "100vh", backgroundColor: "#753ff6", clipPath: "polygon(8% 0, 100% 0, 100% 100%, 0% 100%)" }}>
                            <div style={{
                                margin: "0 50px 0 50px",
                                position: "absolute", left: "50%", top: "50%",
                                transform: "translate(10%, -50%)"
                            }}>
                                <img src="/admin-signup.png" />
                            </div>
                        </Container>
                    </Box>
                </Grid>
            </Grid>
        </Box >
    );
};

export default CreateAdmin;