package service

import (
	"github.com/jinzhu/copier"
	"meadmin/app/system/model"
	"meadmin/app/system/repo"
	"meadmin/app/system/vo"
)

type SystemDept struct {
	repo *repo.SystemDept
}

func NewSystemDept() *SystemDept {
	service := &SystemDept{}
	service.repo = repo.NewSystemDept()
	return service
}

func (this *SystemDept) GetListTree() ([]vo.DeptTree, error) {
	deptList, err := this.repo.Where("status", model.StatusEnable).List()
	if err != nil {
		return nil, err
	}
	var treeList []vo.DeptTree
	for _, dept := range deptList {
		deptTree := vo.DeptTree{
			ID:       dept.ID,
			Label:    dept.Name,
			ParentId: dept.ParentId,
			Value:    dept.ID,
		}
		treeList = append(treeList, deptTree)
	}
	return this.ToTree(treeList, 0), nil
}

func (this *SystemDept) GetModelListTree() ([]vo.DeptModelTree, error) {
	deptList, err := this.repo.Where("status", model.StatusEnable).List()
	if err != nil {
		return nil, err
	}
	var treeList []vo.DeptModelTree
	for _, dept := range deptList {
		var deptTree vo.DeptModelTree
		copier.Copy(&deptTree, &dept)
		treeList = append(treeList, deptTree)
	}
	return this.ToModelTree(treeList, 0), nil
}

func (this *SystemDept) ToModelTree(data []vo.DeptModelTree, parentId uint64) []vo.DeptModelTree {
	var tree []vo.DeptModelTree
	if len(data) == 0 {
		return tree
	}

	for _, value := range data {

		if value.ParentId == parentId {
			child := this.ToModelTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []vo.DeptModelTree{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}

func (this *SystemDept) ToTree(data []vo.DeptTree, parentId uint64) []vo.DeptTree {
	var tree []vo.DeptTree
	if len(data) == 0 {
		return tree
	}
	for _, value := range data {
		if value.ParentId == parentId {
			child := this.ToTree(data, value.ID)
			if len(child) > 0 {
				value.Children = child
			} else {
				value.Children = []vo.DeptTree{}
			}
			tree = append(tree, value)
		}
	}
	return tree
}
