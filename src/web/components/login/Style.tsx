import {
    createStyles,
    Theme,
    makeStyles,
} from "@material-ui/core/styles";

export const useStyles = makeStyles((theme: Theme) =>
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
            color: "black",
            position: "relative"
        }
    }),
);