package model

import "github.com/HitoroOhria/task-executor/domain/value"

// Tasks はタスク集合
type Tasks []*Task

func (ts Tasks) FindByName(name value.TaskName) *Task {
	for _, t := range ts {
		if t.Name == name {
			return t
		}
	}

	return nil
}

func (ts Tasks) FindByFullName(fullName value.FullTaskName) *Task {
	for _, t := range ts {
		if t.FullName == fullName {
			return t
		}
	}

	return nil
}

func (ts Tasks) FindSelected() *Task {
	for _, t := range ts {
		if t.Selected {
			return t
		}
	}

	return nil
}
