import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import UserList from './components/UserList';
import UserForm from './components/UserForm';
import TaskList from './components/TaskList';
import TaskDetail from './components/TaskDetail';
import TaskForm from './components/TaskForm';

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/users" element={<UserList />} />
        <Route path="/users/new" element={<UserForm />} />
        <Route path="/tasks" element={<TaskList />} />
        <Route path="/tasks/:id" element={<TaskDetail />} />
        <Route path="/tasks/new" element={<TaskForm />} /> { }
        <Route path="/tasks/edit/:id" element={<TaskForm />} /> { }
      </Routes>
    </Router>
  );
};

export default App;
