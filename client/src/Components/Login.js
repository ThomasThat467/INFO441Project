import React, { Component } from "react";
import { Form, FormGroup, Label, Input, Button } from "reactstrap";
import { Redirect } from "react-router-dom";

export class LoginForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            email: '',
            pass: '',
            user: null,
          };
        this.signUpCallback = this.props.signUpCallback;
        this.signInCallback = this.props.signInCallback;
    }

    handleSignUp = () => {
        this.signUpCallback(this.state);
    }

    handleSignIn = () => {
        this.signInCallback(this.state);
        Redirect.to('/list');
    }

    handleChange = (event) => {
        let field = event.target.name; //which input
        let value = event.target.value; //what value
    
        let changes = {}; //object to hold changes
        changes[field] = value; //change this field
        this.setState(changes); //update state
    }

    render() {
        return(
            <div>
                <Form>
                    <FormGroup>
                        <Label for="email">Email</Label>
                        <Input
                            onChange={this.handleChange}
                            type="email"
                            name="email"
                            id="email"
                            placeholder="email..."
                            value={this.state.email}
                        />
                    </FormGroup>
                    <FormGroup>
                        <Label for="pass">Password</Label>
                        <Input 
                            onChange={this.handleChange}
                            type="password"
                            name="pass"
                            id="pass"
                            placeholder="password..."
                            value={this.state.pass}
                        />
                    </FormGroup>
                    <Button color="primary" onClick={this.handleSignUp}>Sign up</Button>
                    <Button color="success" onClick={this.handleSignIn}>Log in</Button>
                </Form>
            </div>
        );
    }
}