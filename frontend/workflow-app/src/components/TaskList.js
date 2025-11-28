import React, { useEffect, useState } from 'react';
import { getTasks } from '../api';

const TaskList = () => {
    const [tasks, setTasks] = useState([]);
    const [filters, setFilters] = useState({
        status: '',
        assignee_id: '',
    });

    useEffect(() => {
        const fetchTasks = async () => {
            try {
                const response = await getTasks(filters);
                setTasks(response.data);
            } catch (error) {
                console.error("Error fetching tasks:", error);
            }
        };

        fetchTasks();
    }, [filters]);

    return (
        <div>
            <h1>Tasks</h1>
            <div>
                <label>Status: </label>
                <select onChange={(e) => setFilters({ ...filters, status: e.target.value })}>
                    <option value="">All</option>
                    <option value="In Progress">In Progress</option>
                    <option value="Completed">Completed</option>
                </select>
            </div>
            <div>
                <label>Assignee: </label>
                <input
                    type="text"
                    onChange={(e) => setFilters({ ...filters, assignee_id: e.target.value })}
                    placeholder="Assignee ID"
                />
            </div>
            <ul>
                {tasks.map((task) => (
                    <li key={task.id}>
                        <h3>{task.title}</h3>
                        <p>{task.summary}</p>
                        <p>Status: {task.status}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default TaskList;
