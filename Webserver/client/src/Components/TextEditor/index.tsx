import React, { FC, useState, useEffect } from 'react'
import Typography from '@mui/material/Typography';
import { useParams } from "react-router-dom";

import { StyledTextField } from './styled';
import { ConfigValues } from '../../Pages/Index/types';


export interface TextEditorProps {
    heading: string;
    state: string[];
    mutator: React.Dispatch<React.SetStateAction<any>>;
    setItems: (event: React.ChangeEvent<HTMLTextAreaElement>, mutator: React.Dispatch<React.SetStateAction<any>>) => void
}

export const TextEditor = ({ heading, state, mutator, setItems }: TextEditorProps) => {
    const { id: indexId } = useParams();

    const handleChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
      setItems(event, mutator);
    };

    return (
      <>
            <Typography style={{ marginBottom: '10px' }} variant="h6" component="div">
                { heading }
            </Typography>
            <StyledTextField
              label="Seperated By Newlines"
              multiline 
              value={state.join('\n')}
              rows={8}
              onChange={handleChange}
            />
      </>
    )
  }
