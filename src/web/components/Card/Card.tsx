import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { Typography } from '@material-ui/core';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            flexGrow: 1
        },
        paper: {
            padding: theme.spacing(2),
            textAlign: 'center',
            color: theme.palette.text.secondary,
            '&:hover': {
                border: '1px solid #5A26A7',
                color: '#5A26A7'
            }
        },
        cardItem: {
            color: '#000',
            '&:hover': {
                color: '#5A26A7'
            }
        }
    })
);

// export default function Card() {
export const Card: React.FC = () => {
    const classes = useStyles();

    return (
        <div className={classes.root}>
            <Grid container spacing={3}>
                <Grid item xs>
                    <Paper className={classes.paper}>
                        <Typography variant="h6">Employees</Typography>
                        <Typography variant="h4" className={classes.cardItem}>
                            60
                        </Typography>
                    </Paper>
                </Grid>
                <Grid item xs>
                    <Paper className={classes.paper}>
                        <Typography variant="h6">Integrations</Typography>
                        <Typography variant="h4" className={classes.cardItem}>
                            16
                        </Typography>
                    </Paper>
                </Grid>
                <Grid item xs>
                    <Paper className={classes.paper}>
                        <Typography variant="h6">Accepted Invites</Typography>
                        <Typography variant="h4" className={classes.cardItem}>
                            43
                        </Typography>
                    </Paper>
                </Grid>
                <Grid item xs>
                    <Paper className={classes.paper}>
                        <Typography variant="h6">Pending Invites</Typography>
                        <Typography variant="h4" className={classes.cardItem}>
                            64
                        </Typography>
                    </Paper>
                </Grid>
            </Grid>
        </div>
    );
};
