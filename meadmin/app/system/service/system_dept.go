package service

import (
	"admease/app/system/dto"
	"admease/app/system/model"
	"admease/app/system/repo"
	"github.com/jinzhu/copier"
)

type SystemDept struct {
	repo *repo.SystemDept
}

func NewSystemDept() SystemDept {
	service := SystemDept{}
	service.repo = repo.NewSystemDept()
	return service
}

func (this SystemDept) GetListTree() ([]dto.DeptTreeDTO, error) {
	deptList, err := this.repo.Where("status=?", model.StatusEnable).List()
	if err != nil {
		return nil, err
	}
	var treeList []dto.DeptTreeDTO
	for _, dept := range deptList {
		deptTree := dto.DeptTreeDTO{
			ID:       dept.ID,
			Label:    dept.Name,
			ParentId: dept.ParentId,
			Value:    dept.ID,
		}
		treeList = append(treeList, deptTree)
	}
	return this.ToTree(treeList, 0), nil
}

func (this SystemDept) GetModelListTree() ([]dto.DeptModelTree, error) {
	deptList, err := this.repo.Where("status=?", model.StatusEnable).List()
	if err != nil {
		return nil, err
	}
	var treeList []dto.DeptModelTree
	for _, dept := range deptList {
		var deptTree dto.DeptModelTree
		copier.Copy(&deptTree, &dept)
		treeList = append(treeList, deptTree)
	}
	return this.ToModelTree(treeList, 0), nil
}

func (this SystemDept) ToModelTree(data []dto.DeptModelTree, parentId uint64) []dto.DeptModelTree {
	var tree []dto.DeptModelTree
	if len(data) == 0 {
		return tree
	}

	for _, value := range data {

		if value.ParentId == parentId {
			child := this.ToModelTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []dto.DeptModelTree{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}

func (this SystemDept) ToTree(data []dto.DeptTreeDTO, parentId uint64) []dto.DeptTreeDTO {
	var tree []dto.DeptTreeDTO
	if len(data) == 0 {
		return tree
	}

	for _, value := range data {

		if value.ParentId == parentId {
			child := this.ToTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []dto.DeptTreeDTO{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}
