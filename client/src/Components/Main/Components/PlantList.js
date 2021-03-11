import React, {Component} from 'react';
import {WateringSchedule} from './WateringSchedule.js'


export class PlantList extends Component {
    constructor(props) {
        super(props);

        this.state = {plants:[]};
    }

    componentDidMount() {

        // plantsRef.on('value', (snapshot) => {
        //     let value = snapshot.val();
        //     let plantIds = Object.keys(value);
        //     let plants = plantIds.map((plantId) => {
        //         return {id: plantId, ...value[plantId]}
        //     })
        //     this.setState({plants: plants});
        // });
        
    }
    
    render() {
        
        console.log(this.props.plants.Plants);
        let plantList = this.props.plants.Plants.map((plant) => {
            return <PlantCard  key={plant.plantName} plant={plant}></PlantCard>
        })
        console.log(plantList)
        return (
            <div className="row" id="inventory">{plantList}</div>
        );
    }
}

export class PlantCard extends Component {
    render() {
        let plant = this.props.plant;
        return (
          <div className="cards col-sm-12 col-md-6 col-xl-4">
            <div className="card card-horizontal card-img-top">
              <div className="main-card">
                <img src={plant.photoURL} className="card-img" alt={plant.plantName} />
                <div className="card-body">
                    <h2 className="card-title">{plant.plantName}</h2>
                    <p className="card-text">Watering Schedule</p>
                    <WateringSchedule modifiable={false} schedule={plant.wateringSchedule}></WateringSchedule>
                </div>
              </div>
              
              <div className="card-footer">
                  <p>Last watered: 2020-05-14T17:00:10.859Z</p>
              </div>
            </div>
          </div>
          
        );
    }
}