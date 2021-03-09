import React, {Component} from 'react';
import {Badge} from 'reactstrap'

const weekdays = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];

export class WateringSchedule extends Component {
    constructor(props) {
        super(props);
        this.state = {schedule:[]};
        this.updateSchedule = this.updateSchedule.bind(this);
        this.handleWateringScheduleCallback = this.props.handleWateringSchedule;
        this.onChange = this.props.onChange;
    }

    updateSchedule(day, isActive) {
        console.log("updateSchedule called on ",day, " set to ", isActive)
        console.log(this.state)
        if (!this.state.schedule.includes(day)) {
            this.state.schedule.push(day);
        } else {
            this.state.schedule.splice(this.state.schedule.indexOf(day),1);
        }
        this.handleWateringScheduleCallback(this.state.schedule)
    }

    render() {
        let schedule = this.props.schedule || [];
        let modifiable = this.props.modifiable;
        let wateringSchedule = weekdays.map( (weekday) => {
            let active = false;
            if (schedule.includes(weekday)) {
                active = true;
            }
            return <WateringBadge updateScheduleCallback={this.updateSchedule} key={weekday} isActive={active} modifiable={modifiable} weekday={weekday}></WateringBadge>
        })
        return (
            <div className="row">
                {wateringSchedule}
            </div>
        );
    }
}

export class WateringBadge extends Component {
    constructor(props) {
        super(props);
        this.state = {isActive: props.isActive, day: this.props.weekday};
        this.toggleActive = this.toggleActive.bind(this);
        this.updateScheduleCallback = this.props.updateScheduleCallback;
    }

    toggleActive() {
        if (this.props.modifiable) {
            this.setState({isActive: !this.state.isActive});
            this.updateScheduleCallback(this.state.day, !this.state.isActive);
        }
    }

    render() {
        let weekday = this.props.weekday;
        let color = "secondary"
        if (this.state.isActive) {
            color = "success";
        } 
        return (<Badge onClick={this.toggleActive} color={color}>{weekday}</Badge>);
    }
}