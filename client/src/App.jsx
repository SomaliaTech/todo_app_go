import React, { useEffect, useState } from "react";
import axios from "axios";
import TodoItem from "./componets/TodoItem";
export const BASE_URL = "http://localhost:3000/api";
const App = () => {
  const [open, setOpen] = useState(false);
  const [title, setTile] = useState("");
  const [loading, setLoading] = useState(false);
  const [summery, setSummery] = useState("");
  const [todos, setTodos] = useState([]);
  const handleComplated = async (id) => {
    try {
      await axios.put(BASE_URL + `/update/${id}`);
    } catch (err) {
      console.log(err);
    }
  };

  const handleDelete = async (id) => {
    try {
      await axios.delete(BASE_URL + `/delete/${id}`);
    } catch (err) {
      console.log(err);
    }
  };
  useEffect(() => {
    const fetchdata = async () => {
      try {
        const res = await axios.get(BASE_URL + "/todos");
        setTodos(res.data);
      } catch (err) {
        console.log(err);
      }
    };
    fetchdata();
  }, [open, handleDelete, handleComplated]);

  const handelCreate = async () => {
    try {
      setLoading(true);
      await axios.post(BASE_URL + "/create", { title, summery });
      setLoading(false);
      setOpen(false);
    } catch (err) {
      setLoading(false);
      console.log(err);
    }
  };

  return (
    <div>
      <div className="container">
        <h1 className="title">My Tasks</h1>
        {todos?.length > 0 ? (
          <>
            {todos.map((item) => (
              <TodoItem
                key={item.id}
                handleDelete={handleDelete}
                item={item}
                handleComplated={handleComplated}
              />
            ))}
          </>
        ) : (
          <img src="../public/go.png" className="go_image" />
        )}
        <button onClick={() => setOpen(true)} className="btn">
          New Tasks
        </button>
      </div>
      {/* model */}
      {open && (
        <div className="model">
          <div className="model-box">
            <h1 className="title-box">New Task</h1>
            <div className="item">
              <label>
                Title <span style={{ color: "red", fontSize: 18 }}>*</span>
              </label>
              <input
                onChange={(e) => setTile(e.target.value)}
                type="text"
                placeholder="Title"
                className="input"
              />
            </div>
            <div className="item">
              <label>
                Description<span style={{ color: "red", fontSize: 18 }}>*</span>
              </label>
              <input
                onChange={(e) => setSummery(e.target.value)}
                type="text"
                placeholder="Descriptions"
                className="input"
              />
            </div>
            <div className="icons">
              <button onClick={() => setOpen(false)} className="button cancel">
                Cancel
              </button>
              <button onClick={handelCreate} className="button">
                {loading ? "loading..." : "Create Task"}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default App;
