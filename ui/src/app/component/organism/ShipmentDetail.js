import React, {Component} from 'react';
import {connect} from 'react-redux';
import {getHistoricalLocations} from '../../action';

import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

function mapStateToProps(state) {
  return {
    locationList: state.shipmentDetail.locationList
  };
}
const mapDispatchToProps = (dispatch, ownProps) => ({
  getHistoricalLocations: (shipmentKey) => dispatch(getHistoricalLocations(shipmentKey))
});

class ShipmentDetail extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  componentDidMount() {
    console.log(this.props.params.id)
    this.props.getHistoricalLocations(this.props.params.id);
  }

  render() {
    const locationListArray = this.props.locationList
    let locationCells = ""

    if (locationListArray) {
      locationCells = locationListArray.map((value, key) => {
        console.log(value)
        return (
          <TableRow>
            <TableCell>{value.Record.latitude}</TableCell>
            <TableCell>{value.Record.lognitude}</TableCell>
            <TableCell>{value.Record.locationDatetimeUtc}</TableCell>
          </TableRow>
        )
      })
    }

    return (
      <div className="scdemo">
        <h1>Shipment Key: {this.props.params.id}</h1>
        <h1>Historical Shipments Tracking</h1>
        <Table>
          <TableBody>
            {locationCells}
          </TableBody>
        </Table>
      </div>
    )
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ShipmentDetail);