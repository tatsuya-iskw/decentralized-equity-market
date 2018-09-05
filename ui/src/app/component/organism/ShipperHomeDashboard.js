// DATA
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {fetchShareByUserId} from '../../action';
import {buyShares} from '../../action';
import {scanCorporation} from '../../action';
import {fetchPortfolio} from '../../action';
// UI
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Dialog from '@material-ui/core/Dialog';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import CircularProgress from '@material-ui/core/CircularProgress';
import CheckIcon from '@material-ui/icons/Check';

import Button from '@material-ui/core/Button';

function mapStateToProps(state) {
  return {
    portfolio: state.shipperReducer.portfolio,
    shipmentList: state.shipperReducer.shipmentList,
    shares: state.shipperReducer.shares,
    corp_total: state.shipperReducer.corp_total,
    corp_own: state.shipperReducer.corp_own
  };
}
const mapDispatchToProps = (dispatch, ownProps) => ({
  fetchPortfolio: () => dispatch(fetchPortfolio()),
  fetchAllShare: (shipper_id) => dispatch(fetchShareByUserId(shipper_id)),
  buyShares: () => dispatch(buyShares()),
  scanCorporation: () => dispatch(scanCorporation())
});

class ShipperHomeDashboard extends Component {
  constructor(props) {
    super(props)
    this.state = {
      open: false,
      scanning: false,
      loading: false,
      success: false,
      scaned: false,
      complete: false
    }
  }

  componentWillUnmount() {
    clearTimeout(this.timer);
  }

  componentDidMount() {
    this.props.fetchPortfolio();
    // this.props.fetchAllShare("4F30880E84638F3B4F2DCFE3CC53022CCAE10D491689E0DAC13B85C9C3A8F0E7");
  }

  handleBuyOrder = () => {
    this.props.buyShares()

    this.setState(
      {
        open: true,
        loading: true,
        success: false
      }, () => {
        this.timer = setTimeout(() => {
          this.setState({
            loading: false,
            complete: true,
          });
        }, 8000)
    })
  };

  handleClose = () => {
    this.setState({ open: false });
  };

  handleScanCorporation = () => {
    this.setState({
        scanning: true,
      }, () => {
        this.timer = setTimeout(() => {
          this.props.scanCorporation()
          this.setState({
            scanning: false,
            scaned: true
          });
        }, 6000)
    })
  };

  handleClickOpen = () => {
    this.setState(
      {
        open: true,
        loading: true,
        success: false
      }, () => {
        this.timer = setTimeout(() => {
          this.setState({
            loading: false,
            success: true,
          });
        }, 2000);
      });
  };

  render() {
    console.log(this.props.portfolio)

    const sharesList = this.props.shares
    const portfolioList = this.props.portfolio

    let portfolioCells = ""
    let ShareCells = ""

    if (portfolioList.length > 0) {
      console.log(portfolioList)

      let asset = portfolioList[0].asset
      let number = portfolioList[1].number_share
      let price = portfolioList[2].market_price

      let mv = parseInt(number) * parseInt(price)

      let styles = {}
      if (asset === "ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def") {
        styles = {fontColor: {color: "#2E7D32"}}
      }

      portfolioCells = (
        <TableRow>
          <TableCell style={styles.fontColor}>{asset}</TableCell>
          <TableCell>{number}</TableCell>
          <TableCell>{price}</TableCell>
          <TableCell>{mv}</TableCell>
        </TableRow>
      )

      // portfolioCells = sharesList.map((value, key) => {
      //   let mv = parseInt(value.Record.number_share) * parseInt(value.Record.market_price)
      //   return (
      //     <TableRow>
      //       <TableCell>{value.Record.asset}</TableCell>
      //       <TableCell>{value.Record.number_share}</TableCell>
      //       <TableCell>{value.Record.market_price}</TableCell>
      //       <TableCell>{mv}</TableCell>
      //     </TableRow>
      //   )
      // })
    }

    if (sharesList) {
      ShareCells = sharesList.map((value, key) => {
        return (
          <TableRow>
            <TableCell>{value.Key}</TableCell>
            <TableCell>{value.Record.share_corporate_id}</TableCell>
            <TableCell>{value.Record.share_hash_id}</TableCell>
            <TableCell>{value.Record.share_owner_id}</TableCell>
          </TableRow>
        )
      })
    }

    let dialogTitle = "Cleaning Swap Transaction..."
    if (this.state.success) {
      dialogTitle = "Cleaning Done"
    } else if (this.state.complete) {
      dialogTitle = "Asset Swapping Done"
    }

    return (
      <div className="scdemo">
        <h1>USER EQUITY PORTFOLIO</h1>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>assets</TableCell>
              <TableCell>holding shares</TableCell>
              <TableCell>market price</TableCell>
              <TableCell>market value</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {portfolioCells}
            {/* {ShareCells} */}
            <TableRow>
              <TableCell>cash</TableCell>
              <TableCell>-</TableCell>
              <TableCell>-</TableCell>
              <TableCell>1000</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>total</TableCell>
              <TableCell>-</TableCell>
              <TableCell>-</TableCell>
              <TableCell>1,480</TableCell>
            </TableRow>
          </TableBody>
        </Table>

        <h1>SCAN CORPORATION</h1>

        {this.state.scaned?
          <div>
          <h2>Alice Inc. - <span style={{color: "#2E7D32"}}>ab06a741577bf44e459b8c0c163196727f0e0ae1649db3e05c0cb5a3d6213def</span></h2>
          <h3>Outstanding total shares: {this.props.corp_own.company_owned_shares}</h3>
          <h3>shares Hold by company: {this.props.corp_total.company_owned_shares}</h3>
          <h4>would you buy some shares?</h4>
          <h4>HOW MANY: <input type="text" /></h4>
          <h4>PER PRICE: <input type="text" /></h4>

          <Button onClick={this.handleClickOpen}>BID!</Button>
          </div>
        : <Button onClick={this.handleScanCorporation}>SCAN CORPORATION</Button>}

        {this.state.scanning? <div><CircularProgress size={68} /></div> : null}

        <Dialog
          open={this.state.open}
          onClose={this.handleClose}>
          <DialogTitle id="alert-dialog-title">{dialogTitle}</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              {this.state.loading ? <CircularProgress size={68} /> : null }
              {this.state.success ? 
                <div>
                  <CheckIcon size={68} />
                  <div>would you really swap assets?</div>
                  <div>2 shares</div>
                  <div>20 per shares</div>
                  <div>40 EURO</div>
                  <Button onClick={this.handleBuyOrder}>Make an order</Button>
                </div>
                 : null }
              {this.state.complete ? 
                <div>
                  <Button onClick={this.handleClose}>Success</Button>
                </div>
                 : null }
            </DialogContentText>
          </DialogContent>
        </Dialog>
      </div>
    );
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ShipperHomeDashboard);