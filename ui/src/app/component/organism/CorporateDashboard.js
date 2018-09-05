// Data
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router';
import {fetchCorporateShares} from '../../action';
// UI
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';

import Button from '@material-ui/core/Button';
import Tooltip from '@material-ui/core/Tooltip';

function mapStateToProps(state) {
  return {
    shares: state.corporateVisualization.shares
  };
}
const mapDispatchToProps = (dispatch, ownProps) => ({
  fetchCorporateShares: () => dispatch(fetchCorporateShares())
});

class CorporateDashboard extends Component {
  constructor(props) {
    super(props)
    this.state = {}
  }

  componentDidMount() {
    this.props.fetchCorporateShares();
  }

  render() {
    console.log(this.props.shares)
    const sharesListArray = this.props.shares
    let shareCells = ""

    if (sharesListArray) {
      shareCells = sharesListArray.map((value, key) => {
        let styles = {}

        if (value.Record.share_owner_id === "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def") {
          styles = {
            fColor: {
                color: "#2E7D32"
            }
          }
        } else if (value.Record.share_owner_id === "4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7") {
          styles = {
            fColor: {
                color: "#1565C0"
            }
          }
        }

        return (
          <TableRow>
            <TableCell>
              <Tooltip title={value.Record.share_owner_id} placement="right">
                <Button style={styles.fColor}>{value.Record.share_hash_id}</Button>
              </Tooltip>
            </TableCell>
          </TableRow>
        )
      })
    }

    return (
      <div className="scdemo">
        <h1>Corporate Dashboard</h1>
        <Table>
          <TableBody>
            {shareCells}
          </TableBody>
        </Table>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(CorporateDashboard);