import React, { useState } from 'react';
import PropTypes from 'prop-types';
import api from '../../../../Constants/APIEndpoints';
import Errors from '../../../Errors/Errors';
import { Button } from 'reactstrap';

const SignOutButton = ({ setAuthToken, setUser }) => {
    const [error, setError] = useState("");

    return <>
    <Button className="btn btn-primary" onClick={async (e) => {
        e.preventDefault();
        const response = await fetch(api.base + api.handlers.sessionsMine, {
          method: "DELETE",
          headers: new Headers({
            "Authorization": localStorage.getItem("Authorization")
          })
        });
        if (response.status >= 300) {
            const error2 = await response.text();
            setError(error2);
            return;
        }
        localStorage.removeItem("Authorization");
        setError("");
        setAuthToken("");
        setUser(null);
    }}>Sign out</Button>
        {error &&
            <div>
                <Errors error={error} setError={setError} />
            </div>
        }
    </>
}

SignOutButton.propTypes = {
    setAuthToken: PropTypes.func.isRequired,
    setUser: PropTypes.func.isRequired
}

export default SignOutButton;