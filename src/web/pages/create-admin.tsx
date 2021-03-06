import Box from "@material-ui/core/Box";
import Grid from "@material-ui/core/Grid";
import { Container, FormControl, Button, Typography } from "@material-ui/core";
import { TextField } from "@material-ui/core";
import "@fontsource/poppins";
import InputLabel from '@material-ui/core/InputLabel';
import {
    createStyles,
    fade,
    Theme,
    withStyles,
    makeStyles,
} from '@material-ui/core/styles';
import InputBase from '@material-ui/core/InputBase';

const BootstrapInput = withStyles((theme: Theme) =>
    createStyles({
        root: {
            'label + &': {
                marginTop: theme.spacing(3),
            },
        },
        input: {
            borderRadius: 4,
            position: 'relative',
            backgroundColor: 'rgba(196, 196, 196, 0.2)',
            border: '1px solid black',
            fontSize: 16,
            width: '55vh',
            padding: '10px 12px',
            transition: theme.transitions.create(['border-color', 'box-shadow']),
            // Use the system font instead of the default Roboto font.
            fontFamily: [
                '-apple-system',
                'BlinkMacSystemFont',
                '"Segoe UI"',
                'Roboto',
                '"Helvetica Neue"',
                'Arial',
                'sans-serif',
                '"Apple Color Emoji"',
                '"Segoe UI Emoji"',
                '"Segoe UI Symbol"',
            ].join(','),
            '&:focus': {
                boxShadow: `${fade(theme.palette.primary.main, 0.25)} 0 0 0 0.2rem`,
                borderColor: theme.palette.primary.main,
            },
        },
    }),
)(InputBase);

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            flexWrap: 'wrap',
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
                    {/* <Box>
                        <img src="/admin-signup.png" width="50" height="50" />
                    </Box> */}
                    <Container style={{ height: "100vh", backgroundColor: "white", display: "flex", alignItems: "center", margin: "50px 0 0 0" }} >

                        <Grid item container direction="column" alignItems="center">
                            <Typography style={{ color: "#000000", fontWeight: "bold", fontSize: "30px", letterSpacing: "4px" }}>
                                Create Admin Account
                            </Typography>
                            <Typography style={{ color: "#867979" }}>Fill out  the folowing information  to create an admin account</Typography>
                            <br />
                            <FormControl className={classes.margin}>
                                <InputLabel shrink htmlFor="bootstrap-input" className= {classes.input} >
                                    Company's Name
                                </InputLabel>
                                <BootstrapInput id="bootstrap-input"  placeholder="Wuralab" />
                            </FormControl>
                            <FormControl className={classes.margin}>
                                <InputLabel shrink htmlFor="bootstrap-input" className= {classes.input} >
                                    Company's Email
                                </InputLabel>
                                <BootstrapInput id="bootstrap-input" placeholder="user@wuralab.com" />
                            </FormControl>
                            <FormControl className={classes.margin}>
                                <InputLabel shrink htmlFor="bootstrap-input" className= {classes.input} >
                                    Password
                                </InputLabel>
                                <BootstrapInput id="bootstrap-input"  placeholder="password"/>
                            </FormControl>
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