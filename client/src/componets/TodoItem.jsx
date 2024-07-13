import React from "react";

function TodoItem({ item, handleDelete, handleComplated }) {
  return (
    <div className="todo">
      <div className="todo-top">
        <h4 style={{ textDecoration: item.completed && "line-through" }}>
          {item.title}
        </h4>
        <div className="state">
          <div
            style={{
              backgroundColor: item.completed && "green",
              color: item.completed && "white",
            }}
            onClick={() => handleComplated(item.id)}
            className="proggress"
          >
            {item.completed ? "Completed" : "Inprogress"}
          </div>
          <div onClick={() => handleDelete(item.id)} className="delete">
            Delete
          </div>
        </div>
      </div>
      <p
        style={{ textDecoration: item.completed && "line-through" }}
        className="bottom"
      >
        {item.summery}
      </p>
    </div>
  );
}

export default TodoItem;
