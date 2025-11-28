import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

const TaskDetail = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [task, setTask] = useState(null);
    const [status, setStatus] = useState('');
    const [assignee_id, setAssigneeId] = useState('');

    useEffect(() => {
        const fetchTask = async () => {
            try {
                const response = await getTaskById(id);
                const task = response.data;
                setTask(task);
                setStatus(task.status);
                setAssigneeId(task.assignee_id || '');
            } catch (error) {
                console.error('Error fetching task:', error);
            }
        };

        fetchTask();
    }, [id]);

    const handleUpdate = async () => {
        try {
            await updateTask(id, { status, assignee_id });
            navigate('/tasks'); // Используем navigate вместо history.push
        } catch (error) {
            console.error('Error updating task:', error);
        }
    };

    return task ? (
        <div>
            <h1>{task.title}</h1>
            <p>{task.raw_text}</p>
            <div>
                <label>Status: </label>
                <select value={status} onChange={(e) => setStatus(e.target.value)}>
                    <option value="In Progress">In Progress</option>
                    <option value="Completed">Completed</option>
                </select>
            </div>
            <div>
                <label>Assignee: </label>
                <input
                    type="text"
                    value={assignee_id}
                    onChange={(e) => setAssigneeId(e.target.value)}
                />
            </div>
            <button onClick={handleUpdate}>Update Task</button>
        </div>
    ) : (
        <p>Loading...</p>
    );
};

export default TaskDetail;
