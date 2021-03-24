import {
    createStyles,
    fade,
    Theme,
    withStyles,
} from "@material-ui/core/styles";
import InputBase from "@material-ui/core/InputBase";


export const BootstrapInput = withStyles((theme: Theme) =>
    createStyles({
        root: {
            "label + &": {
                marginTop: theme.spacing(3),
            },
        },
        input: {
            borderRadius: 4,
            position: "relative",
            backgroundColor: "rgba(196, 196, 196, 0.2)",
            border: "1px solid black",
            fontSize: 16,
            width: "55vh",
            padding: "10px 12px",
            transition: theme.transitions.create(["border-color", "box-shadow"]),
            // Use the system font instead of the default Roboto font.
            fontFamily: [
                "-apple-system",
                "BlinkMacSystemFont",
            ].join(","),
            "&:focus": {
                boxShadow: `${fade(theme.palette.primary.main, 0.25)} 0 0 0 0.2rem`,
                borderColor: theme.palette.primary.main,
            },
        },
    }),
)(InputBase);
