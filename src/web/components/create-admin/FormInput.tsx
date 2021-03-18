import { FormControl } from "@material-ui/core";
import InputLabel from "@material-ui/core/InputLabel";
import { BootstrapInput } from "./BootstrapInput";
import { useStyles } from "./Style";

export const FormInput = (props) => {
    const classes = useStyles();

    return (
        <FormControl className={classes.margin}>
            <InputLabel shrink htmlFor="bootstrap-input" className={classes.input} >
                {props.name}
            </InputLabel>
            <BootstrapInput id="bootstrap-input" placeholder={props.placeholder} />
        </FormControl>
    )

}