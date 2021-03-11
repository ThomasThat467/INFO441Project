import React, { Component } from 'react';
import { Checkbox } from '@material-ui/core';
import api from '../../../../Constants/APIEndpoints/APIEndpoints';
import Errors from '../../../Errors/Errors';
import { WateringSchedule } from './WateringSchedule';

class AddPlant extends Component {
    constructor(props) {
        super(props);
        this.state = {
            plantName: '',
            wateringSchedule: [],
            lastWatered: new Date(),
            photoURL: '',
            error: ''
        }
    }

    sendRequest = async (e) => {
        e.preventDefault();
        const { plantName, wateringSchedule, lastWatered, photoURL } = this.state;
        const sendData = { plantName, wateringSchedule, lastWatered, photoURL };
        const response = await fetch(api.base + api.handlers.plants, {
            method: "POST",
            body: JSON.stringify(sendData),
            headers: new Headers({
                "Authorization": localStorage.getItem("Authorization"),
                "Content-Type": "application/json"
            })
        });
        if (response.status >= 300) {
            const error = await response.text();
            console.log(error);
            this.setError(error);
            return;
        }
        alert("Added plant") // TODO make this better by refactoring errors
        // const user = await response.json();
        // this.props.setUser(user);
    }

    setValue = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    addDay = (e) => {
        this.state.wateringSchedule.push(e.target.value);
    }

    setError = (error) => {
        this.setState({ error })
    }

    render() {
        const { firstName, lastName, error } = this.state;
        return <>
            <Errors error={error} setError={this.setError} />
            <div>Give info for new plant:</div>
            <form onSubmit={this.sendRequest}>
                <div>
                    <span>Plant Name: </span>
                    <input name={"plantName"} value={firstName} onChange={this.setValue} />
                </div>
                <div>
                    <span>Days to Water Each Week: </span>
                    <WateringSchedule></WateringSchedule>
                </div>
                <input type="submit" value="Change name" />
            </form>
        </>
    }

}

export default AddPlant;