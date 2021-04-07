import React from 'react';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import { Grid } from '@material-ui/core';
import { Sidebar } from '../../components/common/SideBar/Sidebar';
import { Main } from '../../components/common/Main/Main';

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            backgroundColor: '#F7F8FC'
        }
    })
);
const DashboardView: React.FC = () => {
    const classes = useStyles();
    return (
        <>
            <div className={classes.root}>
                <Grid container spacing={0}>
                    <Sidebar />
                </Grid>
                <Main />
            </div>
        </>
    );
};

export default DashboardView;
