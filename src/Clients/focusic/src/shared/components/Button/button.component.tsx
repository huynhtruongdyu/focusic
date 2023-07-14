import React from 'react';
import './button.component.css';

interface ButtonComponentProps {
  title: string;
  onClick?: () => void;
}

export default function ButtonComponent(props: ButtonComponentProps) {
  return (
    <button onClick={props.onClick} className="app-button">
      {props.title}
    </button>
  );
}
