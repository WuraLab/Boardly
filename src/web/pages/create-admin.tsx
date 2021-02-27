import Box from '@material-ui/core/Box';
import Grid, { GridSpacing } from '@material-ui/core/Grid';
import { Container, FormControl, Button } from '@material-ui/core';
import { InputLabel } from '@material-ui/core';
import { Input } from '@material-ui/core'
import { TextField } from '@material-ui/core'
import { FormHelperText } from "@material-ui/core"


const CreateAdmin = () => {
    return (
        <Box padding={0} margin={0}>
            <Grid container spacing={0}>
                <Grid item xs={6} >
                    <Container style={{ height: "100vh", backgroundColor: "white" }}>
                        <Grid item container direction="column" alignItems="center" style={{ height: "100vh", }}>
                            <div>
                                <h1 style={{ fontFamily: "" }}> Signup to Boardly </h1>
                            </div>
                            <FormControl>
                                <TextField variant="outlined" label="Company's name" style={{ width: "50vh", }} />
                            </FormControl>
                            <br />
                            <FormControl>
                                <TextField variant="outlined" label="Company's Email" style={{ width: "50vh", }} />
                            </FormControl>
                            <br />
                            <FormControl>
                                <TextField variant="outlined" label="Password" style={{ width: "50vh", }} />
                            </FormControl>
                            <br /> <br />
                            <Button variant="contained" size="large" style={{ backgroundColor: "#22AD80", color: "#FFFFFF"}} >
                                Create Account
                            </Button>
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