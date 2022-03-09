import React, { useEffect, useState } from 'react';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Button from '@mui/material/Button';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';

import { TextEditor } from '../../Components/TextEditor';

import {
    useParams
  } from "react-router-dom";

  import { FieldConfigValues } from './types';

function Searches(): React.ReactElement {
    const { id: indexId } = useParams();
    const { search_slug: searchSlug } = useParams();

    const [searchId, setSearchId] = useState<string>();
    const [allowedFacets, setAllowedFacets] = useState<string[]>([]);
    const [displayFields, setDisplayFields] = useState<string[]>([]);
    const [highlightFields, setHighlightFields] = useState<string[]>([]);


    useEffect(() => {
        fetch(`http://localhost/admin/index/${indexId}/configure/search/${searchSlug}`).then(res => res.json()).then(response => {
            const facetResponse: string[] = response[FieldConfigValues.FACET_CONFIG] || [];
            const displayResponse: string[] = response[FieldConfigValues.DISPLAY_CONFIG] || [];
            const highlightResponse: string[] = response[FieldConfigValues.HIGHLIGHT_CONFIG] || [];

            setSearchId(response.Id);
            setAllowedFacets(facetResponse)
            setDisplayFields(displayResponse);
            setHighlightFields(highlightResponse);

        }).catch(err => console.error(err));
    }, []);

    const setItems = (event: React.ChangeEvent<HTMLTextAreaElement>, mutator: React.Dispatch<React.SetStateAction<any>>) => {
        const value = event.target.value;
        mutator(value.split('\n'));
    }

    const saveConfig = () => {
        const saveRequest = {
            config: {
                id: searchId,
                slug: searchSlug,
                display_fields: displayFields,
                highlight_fields: highlightFields,
                allowed_facets: allowedFacets,
            }
        };
        fetch(`http://localhost/admin/index/${indexId}/configure/search`, { method: "POST", body: JSON.stringify(saveRequest), headers: {
            'Content-Type': 'application/json'
        } }).then(res => res.json()).then(response => {
        }).catch(err => console.error(err));
    };

    return (
        <>
            <Typography variant="h5" component="div">
                Configure { searchSlug }
            </Typography>
            <hr />
            <Grid style={{ marginBottom: '20px' }} container spacing={2}>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setDisplayFields} state={displayFields} heading='Displayed Fields' />
                </Grid>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setAllowedFacets} state={allowedFacets} heading='Allowed Facets' />
                </Grid>
            </Grid>
            <Grid style={{ marginBottom: '20px' }} container spacing={2}>
                <Grid item xs={6}>
                    <TextEditor setItems={setItems} mutator={setHighlightFields} state={highlightFields} heading='Highlighted Fields' />
                </Grid>
                <Grid item xs={6}>
                </Grid>
            </Grid>
            <hr />
            <br />
            <Button onClick={saveConfig} variant="contained">Save</Button>
        </>
    );
}

export default Searches;
