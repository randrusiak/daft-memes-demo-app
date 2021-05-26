import React, { useState } from 'react';
import "./App.css";
import { Container, Form, Button } from "react-bootstrap";
import ApiClient from './api';

function App() {

  const api = new ApiClient();
  const [title, setTitle] = useState('');
  const [image, setImage] = useState(null);
  const handleSubmit = async () => {
    const res = await api.getMemes();
    console.log("submit");
    console.log(res);
  };

  const handleInput = event => {
    const target = event.target;
    setTitle(target.value);
  }

  const handleFileInput = event => {
    const file = event.target.files[0];
    setImage(file);
  }

  return (
    <div className="App">
      <Container>
        <div className="App-wrapper">
          <Form>
            <Form.Label>Title:</Form.Label>
            <Form.Control type="text" value={title} onChange={handleInput}/>
            <Form.Label>Image:</Form.Label>
            <Form.Control type="file" onChange={handleFileInput} />
            <Button type="button" onClick={handleSubmit}>
              Submit
            </Button>
          </Form>
        </div>
      </Container>
    </div>
  );
}

export default App;
