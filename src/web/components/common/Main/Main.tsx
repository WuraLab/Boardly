import React from 'react';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Grid';
import Grid from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import { makeStyles, useTheme, Theme, createStyles } from '@material-ui/core/styles';

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: 'flex',
            backgroundColor: '#F7F8FC',
            marginTop: '50px'
        },
        drawer: {
            [theme.breakpoints.up('sm')]: {
                width: drawerWidth,
                flexShrink: 0
            }
        },
        appBar: {
            [theme.breakpoints.up('sm')]: {
                width: `calc(100% - ${drawerWidth}px)`,
                marginLeft: drawerWidth
            }
        },
        menuButton: {
            marginRight: theme.spacing(2),
            [theme.breakpoints.up('sm')]: {
                display: 'none'
            }
        },
        // necessary for content to be below app bar
        toolbar: theme.mixins.toolbar,
        drawerPaper: {
            width: drawerWidth
        },
        content: {
            flexGrow: 1,
            padding: theme.spacing(3)
        }
    })
);

interface Props {
    /**
     * Injected by the documentation to work in an iframe.
     * You won't need it on your project.
     */
    window?: () => Window;
}

export const Main: React.FC = (props: Props) => {
    const { window } = props;
    const classes = useStyles();
    const theme = useTheme();
    const [mobileOpen, setMobileOpen] = React.useState(false);

    const handleDrawerToggle = () => {
        setMobileOpen(!mobileOpen);
    };

    const bull = <span className={classes.bullet}>•</span>;

    return (
        <div className={classes.root}>
            <main className={classes.content}>
                <div className={classes.toolbar} />
                <Grid container spacing={1}>
                    <Card className={classes.root}>
                        <CardContent>
                            <Typography
                                className={classes.title}
                                color="textSecondary"
                                gutterBottom>
                                Word of the Day
                            </Typography>
                            <Typography variant="h5" component="h2">
                                be{bull}nev{bull}o{bull}lent
                            </Typography>
                            <Typography className={classes.pos} color="textSecondary">
                                adjective
                            </Typography>
                            <Typography variant="body2" component="p">
                                well meaning and kindly.
                                <br />
                                {'"a benevolent smile"'}
                            </Typography>
                        </CardContent>
                    </Card>

                    <Card className={classes.root}>
                        <CardContent>
                            <Typography
                                className={classes.title}
                                color="textSecondary"
                                gutterBottom>
                                Word of the Day
                            </Typography>
                            <Typography variant="h5" component="h2">
                                be{bull}nev{bull}o{bull}lent
                            </Typography>
                            <Typography className={classes.pos} color="textSecondary">
                                adjective
                            </Typography>
                            <Typography variant="body2" component="p">
                                well meaning and kindly.
                                <br />
                                {'"a benevolent smile"'}
                            </Typography>
                        </CardContent>
                    </Card>

                    <Card className={classes.root}>
                        <CardContent>
                            <Typography
                                className={classes.title}
                                color="textSecondary"
                                gutterBottom>
                                Word of the Day
                            </Typography>
                            <Typography variant="h5" component="h2">
                                be{bull}nev{bull}o{bull}lent
                            </Typography>
                            <Typography className={classes.pos} color="textSecondary">
                                adjective
                            </Typography>
                            <Typography variant="body2" component="p">
                                well meaning and kindly.
                                <br />
                                {'"a benevolent smile"'}
                            </Typography>
                        </CardContent>
                    </Card>
                    <Card className={classes.root}>
                        <CardContent>
                            <Typography
                                className={classes.title}
                                color="textSecondary"
                                gutterBottom>
                                Word of the Day
                            </Typography>
                            <Typography variant="h5" component="h2">
                                be{bull}nev{bull}o{bull}lent
                            </Typography>
                            <Typography className={classes.pos} color="textSecondary">
                                adjective
                            </Typography>
                            <Typography variant="body2" component="p">
                                well meaning and kindly.
                                <br />
                                {'"a benevolent smile"'}
                            </Typography>
                        </CardContent>
                    </Card>
                </Grid>
            </main>
        </div>
    );
};
