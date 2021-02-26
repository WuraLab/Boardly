import Box from '@material-ui/core/Box';
import Grid, { GridSpacing } from '@material-ui/core/Grid';
import { Container, FormControl } from '@material-ui/core';
import { InputLabel } from '@material-ui/core';
import { Input } from '@material-ui/core'
import { FormHelperText } from "@material-ui/core"
import Paper from '@material-ui/core/Paper';


const CreateAdmin = () => {
    return (
        <Box padding={0} margin={0}>
            <Grid container style={{ backgroundColor: "yellow" }} spacing={0}>
                <Grid item xs={6} >
                    <Container style={{ height: "100vh", backgroundColor: "white" }}>
                        <Grid item container direction="column">
                        <FormControl>
                            <InputLabel>Company's name</InputLabel>
                            <Input></Input>

                        </FormControl>
                        <FormControl>
                            <InputLabel>Company's Email</InputLabel>
                            <Input></Input>

                        </FormControl>
                        <FormControl>
                            <InputLabel>Password</InputLabel>
                            <Input></Input>

                        </FormControl>
                        </Grid>
                        
                    </Container>
                </Grid>
                <Grid item xs={6}>
                    <Box>
                    <Container style={{ height: "100vh", backgroundColor: "#753FF6",  }}>
                        <Box justifyContent="center">
                            <img src="/admin-signup.png" />

                        </Box>
                    </Container>
                    </Box>
                </Grid>
            </Grid>

        </Box>



    )
}

export default CreateAdmin