import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { createTask, updateTask, getTaskById } from '../api';

const TaskForm = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [title, setTitle] = useState('');
    const [raw_text, setRawText] = useState('');
    const [classification, setClassification] = useState('');
    const [summary, setSummary] = useState('');
    const [assignee_id, setAssigneeId] = useState('');

    useEffect(() => {
        if (id) {
            const fetchTask = async () => {
                try {
                    const response = await getTaskById(id);
                    const task = response.data;
                    setTitle(task.title);
                    setRawText(task.raw_text);
                    setClassification(task.classification || '');
                    setSummary(task.summary || '');
                    setAssigneeId(task.assignee_id || '');
                } catch (error) {
                    console.error('Error fetching task:', error);
                }
            };

            fetchTask();
        }
    }, [id]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            if (id) {
                await updateTask(id, { title, raw_text, classification, summary, assignee_id });
            } else {
                await createTask({ title, raw_text, classification, summary, assignee_id });
            }
            navigate('/tasks');
        } catch (error) {
            console.error('Error creating or updating task:', error);
        }
    };

    return (
        <div>
            <h1>{id ? 'Edit Task' : 'Create Task'}</h1>
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Title:</label>
                    <input
                        type="text"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Raw Text:</label>
                    <textarea
                        value={raw_text}
                        onChange={(e) => setRawText(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Classification:</label>
                    <input
                        type="text"
                        value={classification}
                        onChange={(e) => setClassification(e.target.value)}
                    />
                </div>
                <div>
                    <label>Summary:</label>
                    <textarea
                        value={summary}
                        onChange={(e) => setSummary(e.target.value)}
                    />
                </div>
                <div>
                    <label>Assignee ID:</label>
                    <input
                        type="text"
                        value={assignee_id}
                        onChange={(e) => setAssigneeId(e.target.value)}
                    />
                </div>
                <button type="submit">{id ? 'Update Task' : 'Create Task'}</button>
            </form>
        </div>
    );
};

export default TaskForm;
