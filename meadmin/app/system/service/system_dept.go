package service

import (
	"github.com/jinzhu/copier"
	"meadmin/app/system/dto"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
)

type SystemDept struct {
	repo *repo.SystemDept
}

func NewSystemDept() *SystemDept {
	service := &SystemDept{}
	service.repo = repo.NewSystemDept()
	return service
}

func (this *SystemDept) GetListTree() ([]dto.DeptTree, error) {
	deptList, err := this.repo.Where("status", model.StatusEnable).List()
	if err != nil {
		return nil, err
	}
	var treeList []dto.DeptTree
	for _, dept := range deptList {
		deptTree := dto.DeptTree{
			ID:       dept.ID,
			Label:    dept.Name,
			ParentId: dept.ParentId,
			Value:    dept.ID,
		}
		treeList = append(treeList, deptTree)
	}
	return this.ToTree(treeList, 0), nil
}

func (this *SystemDept) GetModelListTree() ([]dto.DeptModelTree, error) {
	deptList, err := this.repo.Where("status", model.StatusEnable).List()
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

func (this *SystemDept) ToModelTree(data []dto.DeptModelTree, parentId uint64) []dto.DeptModelTree {
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

func (this *SystemDept) ToTree(data []dto.DeptTree, parentId uint64) []dto.DeptTree {
	var tree []dto.DeptTree
	if len(data) == 0 {
		return tree
	}
	for _, value := range data {
		if value.ParentId == parentId {
			child := this.ToTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []dto.DeptTree{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}
