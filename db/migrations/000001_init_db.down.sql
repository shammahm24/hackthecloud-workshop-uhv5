DR
BEGIN
    IF EXISTS (SELECT 1 FROM pg_tables WHERE tablename = 'tasks') THEN
        DROP TABLE tasks;
    END IF;
END $$;