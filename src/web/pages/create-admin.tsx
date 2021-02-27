import Box from '@material-ui/core/Box';
import Grid, { GridSpacing } from '@material-ui/core/Grid';
import { Container, FormControl, Button, Typography } from '@material-ui/core';
import { InputLabel } from '@material-ui/core';
import { Input } from '@material-ui/core'
import { TextField, TextFieldProps } from '@material-ui/core'
import { FormHelperText, OutlinedInput } from "@material-ui/core"
import "@fontsource/poppins"

const CreateAdmin = () => {
    return (
        <Box padding={0} margin={0}>
            <Grid container spacing={0}>
                <Grid item xs={6}>
                    {/* <Box>
                        <img src="/admin-signup.png" width="50" height="50" />
                    </Box> */}
                    <Container style={{ height: "100vh", backgroundColor: "white", display: 'flex', alignItems: 'center' }} >

                        <Grid item container direction="column" alignItems="center">
                            <Typography style={{ color: "#000000", fontWeight:"bold", fontSize: "30px", letterSpacing: "4px"}}>
                                Create Admin Account
                            </Typography>
                            <Typography style={{ color: "#867979"}}>Fill out  the folowing information  to create an admin account</Typography>
                            <br />
                            <FormControl>
                                <TextField variant="outlined" label="Company's name" style={{ width: "60vh", }} />
                            </FormControl>
                            <br />
                            <FormControl>
                                <TextField variant="outlined" label="Company's Email" style={{ width: "60vh", }} />
                            </FormControl>
                            <br />
                            <FormControl>
                                <TextField variant="outlined" label="Password" style={{ width: "60vh", }} />
                            </FormControl>
                            <br /> <br />

                            <Button variant="contained" style={{ backgroundColor: "#22AD80", color: "#FFFFFF", width: "50vh" }} >
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
                        <Container style={{ height: "100vh", backgroundColor: "#753FF6", clipPath: "polygon(8% 0, 100% 0, 100% 100%, 0% 100%)" }}>
                            <div style={{
                                margin: "0 50px 0 50px",
                                position: 'absolute', left: '50%', top: '50%',
                                transform: 'translate(10%, -50%)'
                            }}>
                                <img src="/admin-signup.png" />
                            </div>
                        </Container>
                    </Box>
                </Grid>
            </Grid>
        </Box >
    )
}

export default CreateAdmin