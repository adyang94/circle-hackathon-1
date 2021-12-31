import React, {Component} from "react";
import axios from "axios";
import {Card, Header, Form, Input, Icon} from "semantic-ui-react";

let endpoint = "http://localhost:9000";

class ToDoList extends Component {
  constructor(props){
    super(props);

    this.state = {
      tasks: "",
      items: [],
    };
  }

  componentDidMount () {
    this.getTask();
  }

  render () {
    return(
      <div>
        <div className="row">
          <Header className="header" as="h2" color="yellow">
            TO DO LIST
          </Header>
        </div>
      </div>
    );
  }
}

export default ToDoList;