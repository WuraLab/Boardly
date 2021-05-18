/* eslint-disable @typescript-eslint/explicit-function-return-type */
/* eslint-disable @typescript-eslint/explicit-module-boundary-types */
import React from 'react';
import { FormControl } from '@material-ui/core';
import InputLabel from '@material-ui/core/InputLabel';
import { BootstrapInput } from './BootstrapInput';
import { useStyles } from './Style';
import IconButton from '@material-ui/core/IconButton';
import InputAdornment from '@material-ui/core/InputAdornment';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';

interface State {
    showPassword: boolean;
    password: string;
}

export const FormInput = (props) => {
    const classes = useStyles();
    return (
        <FormControl className={classes.margin}>
            <InputLabel shrink htmlFor="bootstrap-input" className={classes.input}>
                {props.name}
            </InputLabel>
            <BootstrapInput
                id="bootstrap-input"
                name={props.name}
                placeholder={props.placeholder}
                onChange={(event) => {
                    console.log(event.target.value);
                }}
            />
        </FormControl>
    );
};

export const FormInputPassword = (props) => {
    const classes = useStyles();
    const [values, setValues] = React.useState<State>({
        showPassword: true,
        password: ''
    });

    const handleChange = (prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
        setValues({ ...values, [prop]: event.target.value });
    };

    const handleClickShowPassword = () => {
        setValues({ ...values, showPassword: !values.showPassword });
    };

    const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };
    return (
        <FormControl className={classes.margin}>
            <InputLabel shrink htmlFor="bootstrap-input" className={classes.input}>
                {props.name}
            </InputLabel>
            <BootstrapInput
                id="bootstrap-input"
                placeholder={props.placeholder}
                type={values.showPassword ? 'text' : 'password'}
                value={values.password}
                onChange={handleChange('password')}
                endAdornment={
                    <InputAdornment position="end">
                        <IconButton
                            style={{ position: 'absolute', marginLeft: '-55px' }}
                            aria-label="toggle password visibility"
                            onClick={handleClickShowPassword}
                            onMouseDown={handleMouseDownPassword}>
                            {values.showPassword ? <Visibility /> : <VisibilityOff />}
                        </IconButton>
                    </InputAdornment>
                }
            />
        </FormControl>
    );
};
