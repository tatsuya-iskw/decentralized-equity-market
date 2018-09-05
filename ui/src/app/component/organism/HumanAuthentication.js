// Data
import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Link} from 'react-router';
// UI
import Dialog from '@material-ui/core/Dialog';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import CircularProgress from '@material-ui/core/CircularProgress';
import CheckIcon from '@material-ui/icons/Check';

import Button from '@material-ui/core/Button';

function mapStateToProps(state) {
  return {};
}
const mapDispatchToProps = (dispatch, ownProps) => ({});

class HumanAuthentication extends Component {
  constructor(props) {
    super(props)

    this.state = {
      open: false,
      loading: false,
      success: false
    }
  }

  componentDidMount() {}

  handleScanHuman = () =>{
    this.setState(
      {
        open: true,
        loading: true,
        success: false
      }, () => {
        this.timer = setTimeout(() => {
          this.setState({
            loading: false,
            success: true
          });
        }, 6000)
      })
  }

  render() {
    let dialogTitle = "Authenticating Human..."

    return (
      <div className="scdemo" style={{textAlign: "center", paddingTop: 200}}>
        <h1>Human Authentication</h1>
        <Button onClick={this.handleScanHuman}>Are you Human?</Button>

        <Dialog
          open={this.state.open}
          onClose={this.handleClose}>
          <DialogTitle id="alert-dialog-title">{dialogTitle}</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              {this.state.loading ? <CircularProgress size={68} /> : null }
              {this.state.success ? 
                <div>
                  <h1><CheckIcon size={68} /> You are Human Bob.</h1>
                  <Link to="/user-home">
                    <Button>
                      Go to Dashboard
                    </Button>
                  </Link>
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
)(HumanAuthentication);