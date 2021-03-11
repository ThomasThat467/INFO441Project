import React from 'react';
import PropTypes from 'prop-types';

const SignForm = ({ setField, submitForm, values, fields }) => {
    return <>
        <form onSubmit={submitForm}>
            {fields.map(d => {
                const { key, name } = d;
                return <div key={key}>
                    <span>{name}: </span>
                    <input className="m-2 mt-3"
                        value={values[key]}
                        name={key}
                        onChange={setField}
                        type={key === "password" || key === "passwordConf" ? "password" : ''}
                    />
                </div>
            })}
            <input className="btn btn-success mt-2 auth-btn" type="submit" value="Submit" />
        </form>
    </>
}

SignForm.propTypes = {
    setField: PropTypes.func.isRequired,
    submitForm: PropTypes.func.isRequired,
    values: PropTypes.shape({
        email: PropTypes.string.isRequired,
        userName: PropTypes.string,
        firstName: PropTypes.string,
        lastName: PropTypes.string,
        password: PropTypes.string.isRequired,
        passwordConf: PropTypes.string
    }),
    fields: PropTypes.arrayOf(PropTypes.shape({
        key: PropTypes.string,
        name: PropTypes.string
    }))
}

export default SignForm;