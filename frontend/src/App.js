import React, { useState, useEffect } from "react";
import "./App.css";
import { Container, Form, Button } from "react-bootstrap";
import ApiClient from "./api";

const api = new ApiClient();

function App() {
  const [title, setTitle] = useState("");
  const [image, setImage] = useState(null);
  const [memes, setMemes] = useState([]);
  const [fileKey, setFileKey] = useState(new Date());

  useEffect(async () => {
    const response = await api.getMemes();
    setMemes(response);
  }, [setMemes]);

  const clearFileInput = () => {
    setFileKey(new Date());
  }

  const handleSubmit = async (e) => {
    const res = await api.addMeme(title, image);
    if (res) {
      setMemes([...memes, res]);
    }
    setTitle("");
    setImage(null);
    clearFileInput();
  };

  const handleInput = (event) => {
    const target = event.target;
    setTitle(target.value);
  };

  const handleFileInput = (event) => {
    const file = event.target.files[0];
    setImage(file);
  };

  const deleteMeme = async id => {
    const res = await api.deleteMeme(id.toString());
    console.log('delete');
    console.log(res);
  }

  const renderMemes = (list) => {
    return (
      <div className="memes-wrapper">
        {list.map((item) => {
          return (
            <div key={item.id} className="meme">
              <div className="cross" onClick={() => deleteMeme(item.id)}>X</div>
              <p>{item.title}</p>
              <img src={item.imagePath} alt={item.title} />
            </div>
          );
        })}
      </div>
    );
  };

  return (
    <div className="App">
      <Container>
        <div className="App-wrapper">
          <Form>
            <Form.Label>Title:</Form.Label>
            <Form.Control type="text" value={title} onChange={handleInput} />
            <Form.Label>Image:</Form.Label>
            <Form.Control key={fileKey} type="file" onChange={handleFileInput} />
            <Button type="button" onClick={handleSubmit}>
              Submit
            </Button>
          </Form>
          {renderMemes(memes)}
        </div>
      </Container>
    </div>
  );
}

export default App;
