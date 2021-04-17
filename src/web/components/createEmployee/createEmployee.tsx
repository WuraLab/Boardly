/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import React from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Modal from '@material-ui/core/Modal';
import Backdrop from '@material-ui/core/Backdrop';
import Fade from '@material-ui/core/Fade';
import IconButton from '@material-ui/core/IconButton';
import AddEmployee from '../../public/addEmployee.svg';
import Button from '@material-ui/core/Button';
import { FormControl, Grid, Typography, Box, Input } from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        modal: {
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center'
        },
        paper: {
            backgroundColor: theme.palette.background.paper,
            border: '2px solid #000',
            boxShadow: theme.shadows[5],
            padding: theme.spacing(2, 4, 3)
        },
        AddEmployeeText: {
            color: '#17093B',
            '&:hover': {
                color: '#5A26A7'
            },
            backgroundColor: '#F7F8FC',
            marginLeft: '10px'
        }
    })
);

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export default function CreateEmployeeModal() {
    const classes = useStyles();
    const [open, setOpen] = React.useState(false);

    // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
    const handleOpen = () => {
        setOpen(true);
    };

    // eslint-disable-next-line @typescript-eslint/explicit-function-return-type
    const handleClose = () => {
        setOpen(false);
    };

    return (
        <div>
            <IconButton onClick={handleOpen} style={{ backgroundColor: '#F7F8FC' }}>
                <img alt="add employee" src={AddEmployee} />
                <Typography variant="h6" className={classes.AddEmployeeText}>
                    Employees
                </Typography>
            </IconButton>
            <Modal
                aria-labelledby="transition-modal-title"
                aria-describedby="transition-modal-description"
                className={classes.modal}
                open={open}
                onClose={handleClose}
                closeAfterTransition
                BackdropComponent={Backdrop}
                BackdropProps={{
                    timeout: 500
                }}>
                <Fade in={open}>
                    <div className={classes.paper}>
                        <div className={classes.root}>
                            <Grid container spacing={3}>
                                <Box className={classes.forgotPassHeader}>
                                    <p>Send Email Invite</p>
                                </Box>

                                <Box m="10%" mt="20%">
                                    <Typography variant="h6" gutterBottom>
                                        <p>Enter E-mail address</p>
                                    </Typography>

                                    <form noValidate autoComplete="off">
                                        <FormControl fullWidth>
                                            <Input
                                                placeholder="user@email.com"
                                                fullWidth
                                                style={{
                                                    border: '0.2px solid black',
                                                    width: '100%',
                                                    borderRadius: '10px',
                                                    backgroundColor: '#F3F3F3',
                                                    padding: '10px'
                                                }}
                                            />
                                        </FormControl>
                                    </form>

                                    <Box mt="10%">
                                        <Box
                                            display="flex"
                                            justifyContent="center"
                                            m={1}
                                            p={1}
                                            bgcolor="#FFFFFF">
                                            <Button
                                                variant="outlined"
                                                style={{
                                                    width: '100%',
                                                    backgroundColor: '#22AD80',
                                                    color: '#ffff',
                                                    fontSize: '12px',
                                                    fontWeight: 'bold',
                                                    padding: '12px',
                                                    borderRadius: '10px'
                                                }}>
                                                Send Invite
                                            </Button>
                                        </Box>
                                    </Box>
                                </Box>
                            </Grid>
                        </div>
                    </div>
                </Fade>
            </Modal>
        </div>
    );
}
